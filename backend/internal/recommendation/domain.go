package recommendation

import (
	"errors"
	"time"
)

type Status string

const (
	StatusRequested        Status = "requested"
	StatusCreated          Status = "created"
	StatusPresented        Status = "presented"
	StatusAccepted         Status = "accepted"
	StatusRejected         Status = "rejected"
	StatusSuperseded       Status = "superseded"
	StatusUnableToComplete Status = "unable_to_complete"
)

type OptionStatus string

const (
	OptionProposed   OptionStatus = "proposed"
	OptionSelected   OptionStatus = "selected"
	OptionRejected   OptionStatus = "rejected"
	OptionSuperseded OptionStatus = "superseded"
)

var ErrNotFound = errors.New("recommendation not found")

type Recommendation struct {
	ID                  string
	UserProfileID       string
	ContextReference    []byte
	Purpose             string
	Rationale           string
	ConfidenceStatement *string
	Status              Status
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type RecommendationOption struct {
	ID               string
	RecommendationID string
	RecipeID         *string
	Position         int32
	Rationale        string
	Status           OptionStatus
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
