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
	"regexp"
	"strings"
	"time"

	"github.com/dapurpintar/backend/internal/model"
)

const defaultBaseURL = "https://api.openai.com/v1"

// OpenAIClient implements Client against any OpenAI-compatible chat
// completions API. Set the base URL via OPENAI_BASE_URL — defaults
// to OpenAI itself. Works with any proxy that speaks the same JSON
// schema (e.g. opencode.ai Go, Together, Groq, OpenRouter, Ollama).
type OpenAIClient struct {
	apiKey  string
	model   string
	baseURL string
	client  *http.Client
	log     *slog.Logger
}

func NewOpenAIClient(apiKey, modelName, baseURL string, log *slog.Logger) *OpenAIClient {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	// strip trailing slash so we can always concat "/chat/completions"
	baseURL = strings.TrimRight(baseURL, "/")
	return &OpenAIClient{
		apiKey:  apiKey,
		model:   modelName,
		baseURL: baseURL,
		client:  &http.Client{},
		log:     log,
	}
}

func (c *OpenAIClient) Name() string { return "openai" }

const systemPrompt = `You are Dapur Pintar, a quiet cooking companion. You compose fresh home-cooking recipes for a real kitchen — no chef-school techniques, no fancy equipment. Given a list of ingredients the user has on hand, plus a small set of reference recipes as style guidance, produce one new recipe.

Bias toward Indonesian and Southeast Asian flavors when the ingredients suggest it, but follow the ingredients. Be specific with measurements (grams, ml, tbsp). Steps should be numbered, with realistic time estimates. Estimate nutrition honestly per serving.

CRITICAL OUTPUT RULES:
- Output ONLY a single valid JSON object. No prose. No markdown. No code fences. No backticks.
- The response must start with { and end with }.
- All field names must be exactly as specified in the schema in the user message.
- All string values must be plain text (no nested markdown like **bold** or *italic*).
- If you cannot fill a field, omit it — do not return null or empty strings for required fields.`

type openaiRequest struct {
	Model          string      `json:"model"`
	Messages       []openaiMsg `json:"messages"`
	ResponseFormat respFormat  `json:"response_format"`
	Temperature    float64     `json:"temperature"`
	MaxTokens      int         `json:"max_tokens"`
}

type openaiMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// respFormat uses json_object mode (the universal "give me JSON" flag).
// OpenAI direct, OpenCode Go, Together, Groq, OpenRouter, Ollama — all
// support it. We previously used strict json_schema but that's only well
// supported on newer OpenAI models; the strict format causes refusals
// or empty responses on most open-source models.
type respFormat struct {
	Type string `json:"type"`
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
		MaxTokens:   3000,
		Messages: []openaiMsg{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		ResponseFormat: respFormat{Type: "json_object"},
	}

	raw, _ := json.Marshal(body)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/chat/completions", bytes.NewReader(raw))
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
		return nil, fmt.Errorf("llm call: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("llm %d: %s", resp.StatusCode, truncate(string(bodyBytes), 400))
	}

	var oa openaiResponse
	if err := json.Unmarshal(bodyBytes, &oa); err != nil {
		return nil, fmt.Errorf("decode llm response: %w", err)
	}
	if oa.Error != nil {
		return nil, fmt.Errorf("llm error: %s", oa.Error.Message)
	}
	if len(oa.Choices) == 0 {
		return nil, errors.New("llm returned no choices")
	}

	recipe, err := parseAndHydrate(oa.Choices[0].Message.Content)
	if err != nil {
		return nil, fmt.Errorf("parse llm output: %w", err)
	}
	return recipe, nil
}

// codeFenceRegex matches ```...``` and ```json ... ``` blocks
var codeFenceRegex = regexp.MustCompile("(?s)```(?:json)?\\s*\\n?(.*?)\\s*```")

// extractJSON pulls the first top-level JSON object out of an LLM
// response. Handles three common failure modes:
//  1. response wrapped in ```json ... ``` fences
//  2. prose preamble ("Here is your recipe: { ... } Thanks!")
//  3. trailing garbage after the closing brace
func extractJSON(raw string) (string, error) {
	s := strings.TrimSpace(raw)

	// First try: strip code fences if present
	if m := codeFenceRegex.FindStringSubmatch(s); len(m) >= 2 {
		return strings.TrimSpace(m[1]), nil
	}

	// Second try: find the first { and the matching last }
	start := strings.Index(s, "{")
	if start < 0 {
		return "", errors.New("no '{' found in llm response")
	}
	end := strings.LastIndex(s, "}")
	if end < 0 || end <= start {
		return "", errors.New("no matching '}' found in llm response")
	}
	return s[start : end+1], nil
}

// parseAndHydrate parses the LLM JSON output, applies field-level
// defaults for anything the model didn't fill in, and hydrates
// presentation-only fields (id, slug, gradient, createdAt).
func parseAndHydrate(raw string) (*model.Recipe, error) {
	clean, err := extractJSON(raw)
	if err != nil {
		return nil, err
	}

	var draft struct {
		Title         string                   `json:"title"`
		TitleID       string                   `json:"titleId"`
		Description   string                   `json:"description"`
		DescriptionID string                   `json:"descriptionId"`
		Cuisine       string                   `json:"cuisine"`
		Difficulty    string                   `json:"difficulty"`
		PrepTime      *int                     `json:"prepTime"`
		CookTime      *int                     `json:"cookTime"`
		Servings      *int                     `json:"servings"`
		Ingredients   []map[string]interface{} `json:"ingredients"`
		Steps         []map[string]interface{} `json:"steps"`
		Nutrition     map[string]interface{}   `json:"nutrition"`
		Tags          []string                 `json:"tags"`
		Dietary       []string                 `json:"dietary"`
	}
	if err := json.Unmarshal([]byte(clean), &draft); err != nil {
		return nil, fmt.Errorf("json unmarshal: %w (clean: %s)", err, truncate(clean, 200))
	}

	if draft.Title == "" {
		return nil, errors.New("llm response missing required field: title")
	}
	if draft.TitleID == "" {
		draft.TitleID = draft.Title
	}
	if draft.Description == "" {
		draft.Description = "A home-cooked recipe composed for your kitchen."
	}
	if draft.DescriptionID == "" {
		draft.DescriptionID = "Resep rumahan yang disusun untuk dapur Anda."
	}
	if draft.Cuisine == "" {
		draft.Cuisine = "Indonesian"
	}
	if draft.Difficulty == "" {
		draft.Difficulty = string(model.DifficultyMedium)
	}
	if draft.PrepTime == nil || *draft.PrepTime <= 0 {
		draft.PrepTime = intPtr(10)
	}
	if draft.CookTime == nil || *draft.CookTime < 0 {
		draft.CookTime = intPtr(15)
	}
	if draft.Servings == nil || *draft.Servings <= 0 {
		draft.Servings = intPtr(2)
	}
	if len(draft.Ingredients) == 0 {
		draft.Ingredients = []map[string]interface{}{minimalIngredient()}
	}
	if len(draft.Steps) == 0 {
		draft.Steps = []map[string]interface{}{minimalStep(1)}
	}

	recipe := &model.Recipe{
		ID:           "gen_" + newID(),
		Slug:         slugify(draft.Title),
		Title:        strings.TrimSpace(draft.Title),
		TitleID:      strings.TrimSpace(draft.TitleID),
		Description:  strings.TrimSpace(draft.Description),
		DescriptionID: strings.TrimSpace(draft.DescriptionID),
		Image:        "",
		Gradient:     []string{"#A8553A", "#723627"},
		Cuisine:      draft.Cuisine,
		Difficulty:   model.Difficulty(draft.Difficulty),
		PrepTime:     *draft.PrepTime,
		CookTime:     *draft.CookTime,
		Servings:     *draft.Servings,
		Ingredients:  normalizeIngredients(draft.Ingredients),
		Steps:        normalizeSteps(draft.Steps),
		Nutrition:    normalizeNutrition(draft.Nutrition),
		Tags:         defaultSlice(draft.Tags),
		Dietary:      normalizeDietary(draft.Dietary),
		Rating:       0,
		Reviews:      0,
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}
	return recipe, nil
}

func intPtr(i int) *int { return &i }

func minimalIngredient() map[string]interface{} {
	return map[string]interface{}{
		"name":     "Ingredients to taste",
		"nameId":   "Bahan secukupnya",
		"amount":   "as needed",
		"category": "other",
	}
}

func minimalStep(order int) map[string]interface{} {
	return map[string]interface{}{
		"order":  order,
		"text":   "Combine ingredients and cook until done.",
		"textId": "Campur bahan dan masak hingga matang.",
	}
}

func normalizeIngredients(raw []map[string]interface{}) []model.RecipeIngredient {
	out := make([]model.RecipeIngredient, 0, len(raw))
	for _, m := range raw {
		name, _ := m["name"].(string)
		nameID, _ := m["nameId"].(string)
		if nameID == "" {
			nameID = name
		}
		amount, _ := m["amount"].(string)
		if amount == "" {
			amount = "to taste"
		}
		category, _ := m["category"].(string)
		if !isValidCategory(category) {
			category = "other"
		}
		optional, _ := m["optional"].(bool)
		out = append(out, model.RecipeIngredient{
			Name:     strings.TrimSpace(name),
			NameID:   strings.TrimSpace(nameID),
			Amount:   amount,
			Category: model.Category(category),
			Optional: optional,
		})
	}
	if len(out) == 0 {
		out = append(out, model.RecipeIngredient{
			Name: "Ingredients to taste", NameID: "Bahan secukupnya",
			Amount: "as needed", Category: model.CategoryOther,
		})
	}
	return out
}

func normalizeSteps(raw []map[string]interface{}) []model.RecipeStep {
	out := make([]model.RecipeStep, 0, len(raw))
	for i, m := range raw {
		var order int
		switch v := m["order"].(type) {
		case float64:
			order = int(v)
		case int:
			order = v
		default:
			order = i + 1
		}
		text, _ := m["text"].(string)
		textID, _ := m["textId"].(string)
		if textID == "" {
			textID = text
		}
		var durationSec *int
		if v, ok := m["durationSec"].(float64); ok && v > 0 {
			d := int(v)
			durationSec = &d
		} else if v, ok := m["durationSec"].(int); ok && v > 0 {
			durationSec = &v
		}
		tip, _ := m["tip"].(string)
		out = append(out, model.RecipeStep{
			Order:       order,
			Text:        strings.TrimSpace(text),
			TextID:      strings.TrimSpace(textID),
			DurationSec: durationSec,
			Tip:         strings.TrimSpace(tip),
		})
	}
	return out
}

func normalizeNutrition(raw map[string]interface{}) model.Nutrition {
	get := func(key string, def int) int {
		if v, ok := raw[key]; ok {
			switch n := v.(type) {
			case float64:
				return int(n)
			case int:
				return n
			}
		}
		return def
	}
	return model.Nutrition{
		Calories: get("calories", 0),
		Protein:  get("protein", 0),
		Carbs:    get("carbs", 0),
		Fat:      get("fat", 0),
		Fiber:    get("fiber", 0),
	}
}

func normalizeDietary(raw []string) []model.DietaryTag {
	out := make([]model.DietaryTag, 0, len(raw))
	for _, s := range raw {
		if isValidDietary(s) {
			out = append(out, model.DietaryTag(s))
		}
	}
	return out
}

func defaultSlice(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}

func isValidCategory(c string) bool {
	switch model.Category(c) {
	case model.CategoryProtein, model.CategoryVegetable, model.CategorySpice,
		model.CategoryGrain, model.CategoryDairy, model.CategorySauce, model.CategoryOther:
		return true
	}
	return false
}

func isValidDietary(d string) bool {
	switch model.DietaryTag(d) {
	case model.DietaryVegetarian, model.DietaryVegan, model.DietaryHalal,
		model.DietaryGlutenFree, model.DietarySpicy, model.DietaryLowCarb:
		return true
	}
	return false
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

	b.WriteString(`
Return a single JSON object with EXACTLY these fields (and nothing else at the top level):
{
  "title":          string,  // recipe name, 2-80 chars
  "titleId":        string,  // same in the output language above
  "description":    string,  // 1-2 sentences, 10-240 chars
  "descriptionId":  string,  // same in the output language
  "cuisine":        string,  // e.g. "Indonesian", "Italian"
  "difficulty":     "easy" | "medium" | "hard"
  "prepTime":       integer,  // minutes, 1-240
  "cookTime":       integer,  // minutes, 0-600
  "servings":       integer,  // 1-24
  "ingredients": [
    {
      "name":     string,         // ingredient, English
      "nameId":   string,         // same in output language
      "amount":   string,         // e.g. "2 tbsp", "150g"
      "category": "protein" | "vegetable" | "spice" | "grain" | "dairy" | "sauce" | "other",
      "optional": boolean         // true if garnish/optional
    }
    // 3-24 items
  ],
  "steps": [
    {
      "order":       integer,  // 1-indexed
      "text":        string,   // step, English, 10+ chars
      "textId":      string,   // same in output language
      "durationSec": integer,  // optional, 0+
      "tip":         string    // optional chef tip
    }
    // 3-16 items
  ],
  "nutrition": {
    "calories": integer,  // per serving
    "protein":  integer,  // grams
    "carbs":    integer,  // grams
    "fat":      integer,  // grams
    "fiber":    integer   // grams
  },
  "tags":    [string],     // 0-12 items, free-form
  "dietary": ["vegetarian" | "vegan" | "halal" | "gluten-free" | "spicy" | "low-carb"]
}

`)

	b.WriteString("Reference recipes (style guidance — do not copy verbatim, use as flavor/structure reference):\n")
	for i, r := range req.References {
		fmt.Fprintf(&b, "\n%d. %s (%s) — cuisine: %s, difficulty: %s, time: %dm+%dm, ingredients: %d, steps: %d\n",
			i+1, r.Title, r.TitleID, r.Cuisine, r.Difficulty, r.PrepTime, r.CookTime,
			len(r.Ingredients), len(r.Steps))
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
	ErrTimeout   = errors.New("llm request timed out")
	ErrRateLimit  = errors.New("llm rate limit")
	ErrBadOutput  = errors.New("llm returned malformed output")
)

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
