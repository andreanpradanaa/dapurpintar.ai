package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/dapurpintar/backend/internal/model"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

// OpenAIClient implements Client against the OpenAI Chat Completions API.
type OpenAIClient struct {
	apiKey string
	model  string
	client *http.Client
	log    *slog.Logger
}

func NewOpenAIClient(apiKey, modelName string, log *slog.Logger) *OpenAIClient {
	return &OpenAIClient{
		apiKey: apiKey,
		model:  modelName,
		client: &http.Client{},
		log:    log,
	}
}

func (c *OpenAIClient) Name() string { return "openai" }

// recipeJSONSchema is the OpenAI structured-output schema for a generated
// recipe. It omits presentation-only fields (gradient, image, rating,
// reviews, createdAt) — those are hydrated by the server after the LLM
// returns. It also omits id/slug since the LLM doesn't know them.
const recipeJSONSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": [
    "title", "titleId", "description", "descriptionId",
    "cuisine", "difficulty", "prepTime", "cookTime", "servings",
    "ingredients", "steps", "nutrition", "tags", "dietary"
  ],
  "properties": {
    "title":          { "type": "string", "minLength": 2, "maxLength": 80 },
    "titleId":        { "type": "string", "minLength": 2, "maxLength": 80 },
    "description":    { "type": "string", "minLength": 10, "maxLength": 240 },
    "descriptionId":  { "type": "string", "minLength": 10, "maxLength": 240 },
    "cuisine":        { "type": "string", "minLength": 2, "maxLength": 40 },
    "difficulty":     { "type": "string", "enum": ["easy","medium","hard"] },
    "prepTime":       { "type": "integer", "minimum": 1, "maximum": 240 },
    "cookTime":       { "type": "integer", "minimum": 0, "maximum": 600 },
    "servings":       { "type": "integer", "minimum": 1, "maximum": 24 },
    "ingredients": {
      "type": "array",
      "minItems": 3,
      "maxItems": 24,
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["name","nameId","amount","category"],
        "properties": {
          "name":     { "type": "string", "minLength": 1 },
          "nameId":   { "type": "string", "minLength": 1 },
          "amount":   { "type": "string", "minLength": 1 },
          "category": { "type": "string", "enum": ["protein","vegetable","spice","grain","dairy","sauce","other"] },
          "optional": { "type": "boolean" }
        }
      }
    },
    "steps": {
      "type": "array",
      "minItems": 3,
      "maxItems": 16,
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["order","text","textId"],
        "properties": {
          "order":       { "type": "integer", "minimum": 1 },
          "text":        { "type": "string", "minLength": 10 },
          "textId":      { "type": "string", "minLength": 10 },
          "durationSec": { "type": "integer", "minimum": 0 },
          "tip":         { "type": "string" }
        }
      }
    },
    "nutrition": {
      "type": "object",
      "additionalProperties": false,
      "required": ["calories","protein","carbs","fat","fiber"],
      "properties": {
        "calories": { "type": "integer", "minimum": 0, "maximum": 5000 },
        "protein":  { "type": "integer", "minimum": 0, "maximum": 200 },
        "carbs":    { "type": "integer", "minimum": 0, "maximum": 500 },
        "fat":      { "type": "integer", "minimum": 0, "maximum": 300 },
        "fiber":    { "type": "integer", "minimum": 0, "maximum": 100 }
      }
    },
    "tags":    { "type": "array", "maxItems": 12, "items": { "type": "string" } },
    "dietary": {
      "type": "array",
      "items": { "type": "string", "enum": ["vegetarian","vegan","halal","gluten-free","spicy","low-carb"] }
    }
  }
}`

const systemPrompt = `You are Dapur Pintar, a quiet cooking companion. You compose fresh home-cooking recipes for a real kitchen — no chef-school techniques, no fancy equipment. Given a list of ingredients the user has on hand, plus a small set of reference recipes as style guidance, produce one new recipe.

Bias toward Indonesian and Southeast Asian flavors when the ingredients suggest it, but follow the ingredients. Be specific with measurements (grams, ml, tbsp). Steps should be numbered, with realistic time estimates. Estimate nutrition honestly per serving.

Output strictly as JSON matching the provided schema. No prose, no markdown, no commentary.`

type openaiRequest struct {
	Model          string         `json:"model"`
	Messages       []openaiMsg    `json:"messages"`
	ResponseFormat respFormat     `json:"response_format"`
	Temperature    float64        `json:"temperature"`
	MaxTokens      int            `json:"max_tokens"`
}

type openaiMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type respFormat struct {
	Type   string         `json:"type"`
	JSONSchema *jsonSchema `json:"json_schema,omitempty"`
}

type jsonSchema struct {
	Name   string `json:"name"`
	Strict bool   `json:"strict"`
	Schema string `json:"schema"`
}

type openaiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

func (c *OpenAIClient) GenerateRecipe(ctx context.Context, req GenerateRequest) (*model.Recipe, error) {
	timeout := req.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	userPrompt := buildUserPrompt(req)

	body := openaiRequest{
		Model:       c.model,
		Temperature: 0.7,
		MaxTokens:   2400,
		Messages: []openaiMsg{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		ResponseFormat: respFormat{
			Type: "json_schema",
			JSONSchema: &jsonSchema{
				Name:   "recipe",
				Strict: true,
				Schema: recipeJSONSchema,
			},
		},
	}

	raw, _ := json.Marshal(body)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", openaiURL, bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(httpReq)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, ErrTimeout
		}
		return nil, fmt.Errorf("openai call: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("openai %d: %s", resp.StatusCode, truncate(string(bodyBytes), 400))
	}

	var oa openaiResponse
	if err := json.Unmarshal(bodyBytes, &oa); err != nil {
		return nil, fmt.Errorf("decode openai response: %w", err)
	}
	if oa.Error != nil {
		return nil, fmt.Errorf("openai error: %s", oa.Error.Message)
	}
	if len(oa.Choices) == 0 {
		return nil, errors.New("openai returned no choices")
	}

	raw2 := oa.Choices[0].Message.Content
	recipe, err := parseAndHydrate(raw2)
	if err != nil {
		return nil, fmt.Errorf("parse LLM output: %w", err)
	}
	return recipe, nil
}

// parseAndHydrate parses the LLM JSON output and fills in
// presentation-only fields (id, slug, gradient, createdAt).
func parseAndHydrate(raw string) (*model.Recipe, error) {
	var draft struct {
		Title        string             `json:"title"`
		TitleID      string             `json:"titleId"`
		Description  string             `json:"description"`
		DescriptionID string             `json:"descriptionId"`
		Cuisine      string             `json:"cuisine"`
		Difficulty   model.Difficulty   `json:"difficulty"`
		PrepTime     int                `json:"prepTime"`
		CookTime     int                `json:"cookTime"`
		Servings     int                `json:"servings"`
		Ingredients  []model.RecipeIngredient `json:"ingredients"`
		Steps        []model.RecipeStep `json:"steps"`
		Nutrition    model.Nutrition    `json:"nutrition"`
		Tags         []string           `json:"tags"`
		Dietary      []model.DietaryTag `json:"dietary"`
	}
	if err := json.Unmarshal([]byte(raw), &draft); err != nil {
		return nil, err
	}

	now := time.Now().UTC().Format(time.RFC3339)
	rec := &model.Recipe{
		ID:           "gen_" + newID(),
		Slug:         slugify(draft.Title),
		Title:        draft.Title,
		TitleID:      draft.TitleID,
		Description:  draft.Description,
		DescriptionID: draft.DescriptionID,
		Image:        "",
		Gradient:     []string{"#A8553A", "#723627"},
		Cuisine:      draft.Cuisine,
		Difficulty:   draft.Difficulty,
		PrepTime:     draft.PrepTime,
		CookTime:     draft.CookTime,
		Servings:     draft.Servings,
		Ingredients:  draft.Ingredients,
		Steps:        draft.Steps,
		Nutrition:    draft.Nutrition,
		Tags:         draft.Tags,
		Dietary:      draft.Dietary,
		Rating:       0,
		Reviews:      0,
		CreatedAt:    now,
	}
	return rec, nil
}

func buildUserPrompt(req GenerateRequest) string {
	var b strings.Builder
	b.WriteString("Language for output: ")
	if req.Language == "id" {
		b.WriteString("Indonesian (Bahasa Indonesia). Use Bahasa Indonesia for all text and titleId fields.\n")
	} else {
		b.WriteString("English. Use English for all text and titleId fields.\n")
	}
	b.WriteString("\nUser has on hand: ")
	b.WriteString(strings.Join(req.UserIngredients, ", "))
	b.WriteString("\n")

	if len(req.Dietary) > 0 {
		b.WriteString("User preferences: ")
		b.WriteString(strings.Join(req.Dietary, ", "))
		b.WriteString("\n")
	}

	b.WriteString("\nReference recipes (style guidance — do not copy verbatim, use as flavor/structure reference):\n")
	for i, r := range req.References {
		fmt.Fprintf(&b, "\n%d. %s (%s) — cuisine: %s, difficulty: %s, time: %dm+%dm, ingredients: %d, steps: %d\n",
			i+1, r.Title, r.TitleID, r.Cuisine, r.Difficulty, r.PrepTime, r.CookTime,
			len(r.Ingredients), len(r.Steps))
		// Sample a few key ingredients
		max := 6
		if len(r.Ingredients) < max {
			max = len(r.Ingredients)
		}
		names := make([]string, 0, max)
		for _, ing := range r.Ingredients[:max] {
			names = append(names, ing.Name+" ("+ing.Amount+")")
		}
		b.WriteString("   key ingredients: ")
		b.WriteString(strings.Join(names, ", "))
		b.WriteString("\n")
	}
	return b.String()
}

var (
	ErrTimeout  = errors.New("openai request timed out")
	ErrRateLimit = errors.New("openai rate limit")
)

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
