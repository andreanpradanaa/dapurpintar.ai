package recommendation

import (
	"context"
	"encoding/json"
	"time"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/ai"
	apperr "github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/errors"
)

type Service struct {
	store   Store
	gateway ai.Gateway
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func NewServiceWithAI(store Store, gateway ai.Gateway) *Service {
	return &Service{store: store, gateway: gateway}
}

const maxPageLimit = 100

type PageInfo struct {
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}

func (s *Service) Request(ctx context.Context, profileID, purpose string, maxPrep *int32, useExpiringFirst bool) (*Recommendation, error) {
	contextRef := []byte("{}")
	rec, err := s.store.CreateRecommendation(ctx, profileID, contextRef, purpose)
	if err != nil {
		return nil, err
	}

	if s.gateway == nil {
		return rec, nil
	}

	profile := ai.DefaultProfile()
	req := ai.Request{
		Purpose:  ai.Purpose(purpose),
		Messages: []ai.Message{{Role: ai.RoleUser, Content: "Sarankan resep masakan untuk hari ini."}},
		Profile:  profile,
	}
	result, aiErr := s.gateway.Complete(ctx, req)
	if aiErr != nil {
		s.store.UpdateRecommendationStatus(ctx, rec.ID, StatusUnableToComplete)
		return rec, nil
	}

	var output struct {
		Summary string `json:"summary"`
		Options []struct {
			Title     string `json:"title"`
			Rationale string `json:"rationale"`
		} `json:"options"`
	}
	if err := json.Unmarshal(result.Content, &output); err != nil {
		s.store.UpdateRecommendationStatus(ctx, rec.ID, StatusUnableToComplete)
		return rec, nil
	}

	for i, opt := range output.Options {
		s.store.CreateRecommendationOption(ctx, rec.ID, nil, int32(i+1), opt.Rationale)
	}

	updated, err := s.store.UpdateRecommendationStatus(ctx, rec.ID, StatusCreated)
	if err == nil {
		rec = updated
	}

	return rec, nil
}

func (s *Service) AnalyzePantry(ctx context.Context, profileID string, items []PantryItemSnapshot) (*PantryAnalysis, error) {
	if s.gateway == nil {
		return &PantryAnalysis{UseFirst: []PantryUseFirst{}, Suggestions: []PantrySuggestion{}}, nil
	}

	contextJSON, _ := json.Marshal(items)
	req := ai.Request{
		Purpose:  ai.PurposePantryAnalysis,
		Messages: []ai.Message{{Role: ai.RoleSystem, Content: "Analyze this pantry and suggest what to use first:\n" + string(contextJSON)}},
		Profile:  ai.DefaultProfile(),
	}
	result, err := s.gateway.Complete(ctx, req)
	if err != nil {
		return &PantryAnalysis{UseFirst: []PantryUseFirst{}, Suggestions: []PantrySuggestion{}}, nil
	}

	var output struct {
		UseFirst    []PantryUseFirst   `json:"use_first_opportunities"`
		Suggestions []PantrySuggestion `json:"optimization_suggestions"`
	}
	if err := json.Unmarshal(result.Content, &output); err != nil {
		return &PantryAnalysis{UseFirst: []PantryUseFirst{}, Suggestions: []PantrySuggestion{}}, nil
	}

	return &PantryAnalysis{UseFirst: output.UseFirst, Suggestions: output.Suggestions}, nil
}

type PantryItemSnapshot struct {
	ID             string  `json:"id"`
	IngredientName string  `json:"ingredient_name"`
	Category       string  `json:"category"`
	Quantity       float64 `json:"quantity"`
	Unit           string  `json:"unit"`
	ExpiryDate     string  `json:"expiry_date,omitempty"`
}

type PantryUseFirst struct {
	PantryItemID   string `json:"pantry_item_id"`
	IngredientName string `json:"ingredient_name"`
	Reason         string `json:"reason"`
}

type PantrySuggestion struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PantryAnalysis struct {
	UseFirst    []PantryUseFirst   `json:"use_first_opportunities"`
	Suggestions []PantrySuggestion `json:"optimization_suggestions"`
}

func (s *Service) Get(ctx context.Context, id string) (*Recommendation, []RecommendationOption, error) {
	rec, err := s.store.GetRecommendationByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	options, err := s.store.ListRecommendationOptions(ctx, id)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}
	return rec, options, nil
}

func (s *Service) List(ctx context.Context, profileID, cursor string, limit int32, status, purpose *string) ([]Recommendation, *PageInfo, error) {
	if limit <= 0 || limit > maxPageLimit {
		limit = 20
	}
	items, err := s.store.ListRecommendations(ctx, profileID, cursor, limit+1, status, purpose)
	if err != nil {
		return nil, nil, apperr.Internal(err)
	}
	hasMore := len(items) > int(limit)
	if hasMore {
		items = items[:limit]
	}
	var nc string
	if len(items) > 0 {
		nc = items[len(items)-1].ID
	}
	return items, &PageInfo{NextCursor: nc, HasMore: hasMore}, nil
}

func (s *Service) Present(ctx context.Context, id string) (*Recommendation, error) {
	rec, err := s.store.GetRecommendationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if rec.Status != StatusCreated {
		return nil, apperr.New(apperr.CodeRecommendationStateInvalid, "Only created recommendations can be presented.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeRecommendationStateInvalid), Message: "The recommendation is not in a presentable state."})
	}
	return s.store.UpdateRecommendationStatus(ctx, id, StatusPresented)
}

func (s *Service) AcceptOption(ctx context.Context, recID, optionID string) (*Recommendation, error) {
	rec, err := s.store.GetRecommendationByID(ctx, recID)
	if err != nil {
		return nil, err
	}
	if rec.Status != StatusPresented {
		return nil, apperr.New(apperr.CodeRecommendationStateInvalid, "The recommendation is not available for acceptance.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeRecommendationStateInvalid), Message: "Only presented recommendations can accept options."})
	}
	opt, err := s.store.GetRecommendationOptionByID(ctx, optionID)
	if err != nil {
		return nil, err
	}
	if opt.Status != OptionProposed {
		return nil, apperr.New(apperr.CodeRecommendationOptionNotOK, "This option cannot be accepted in its current state.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeRecommendationOptionNotOK), Message: "Only proposed options can be accepted."})
	}
	if _, err := s.store.UpdateRecommendationOptionStatus(ctx, optionID, OptionSelected); err != nil {
		return nil, err
	}
	return s.store.UpdateRecommendationStatus(ctx, recID, StatusAccepted)
}

func (s *Service) Reject(ctx context.Context, id string) (*Recommendation, error) {
	rec, err := s.store.GetRecommendationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if rec.Status != StatusPresented && rec.Status != StatusCreated {
		return nil, apperr.New(apperr.CodeRecommendationStateInvalid, "The recommendation cannot be rejected in its current state.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeRecommendationStateInvalid), Message: "The recommendation is not in a rejectable state."})
	}
	return s.store.UpdateRecommendationStatus(ctx, id, StatusRejected)
}

func (s *Service) Supersede(ctx context.Context, id string) (*Recommendation, error) {
	rec, err := s.store.GetRecommendationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if rec.Status == StatusRejected || rec.Status == StatusSuperseded {
		return nil, apperr.New(apperr.CodeRecommendationStateInvalid, "The recommendation cannot be superseded from its current state.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeRecommendationStateInvalid), Message: "The recommendation cannot be superseded."})
	}
	return s.store.UpdateRecommendationStatus(ctx, id, StatusSuperseded)
}

func (s *Service) StartConversation(ctx context.Context, recID string) (*Conversation, error) {
	if _, err := s.store.GetRecommendationByID(ctx, recID); err != nil {
		return nil, err
	}
	conv, err := s.store.GetConversationByRecommendation(ctx, recID)
	if err == nil {
		return conv, nil
	}
	return s.store.CreateConversation(ctx, recID)
}

func (s *Service) GetConversation(ctx context.Context, recID string) (*Conversation, error) {
	return s.store.GetConversationByRecommendation(ctx, recID)
}

func (s *Service) ContinueConversation(ctx context.Context, recID, message string) (*Conversation, error) {
	conv, err := s.store.GetConversationByRecommendation(ctx, recID)
	if err != nil {
		return nil, err
	}
	if conv.Status != "open" {
		return nil, apperr.New(apperr.CodeConversationStateInvalid, "Conversation is not open.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeConversationStateInvalid), Message: "Cannot continue a closed conversation."})
	}

	now := time.Now().UTC()
	conv.Messages = append(conv.Messages, ConversationMessage{Role: "user", Content: message, CreatedAt: now})
	conv.Messages = append(conv.Messages, ConversationMessage{Role: "assistant", Content: "Terima kasih atas pertanyaan Anda. Saya akan membantu.", CreatedAt: now})

	snapshot, err := json.Marshal(conv.Messages)
	if err != nil {
		return nil, apperr.Internal(err)
	}
	return s.store.UpdateConversationSnapshot(ctx, recID, snapshot)
}

func (s *Service) CloseConversation(ctx context.Context, recID string) (*Conversation, error) {
	conv, err := s.store.GetConversationByRecommendation(ctx, recID)
	if err != nil {
		return nil, err
	}
	if conv.Status != "open" {
		return nil, apperr.New(apperr.CodeConversationStateInvalid, "Conversation is not open.").
			WithDetails(apperr.Detail{Field: "status", Code: string(apperr.CodeConversationStateInvalid), Message: "Conversation is already closed."})
	}
	return s.store.CloseConversation(ctx, recID)
}
