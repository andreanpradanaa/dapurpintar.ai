package repo

import (
	"context"
	"errors"

	"github.com/dapurpintar/backend/internal/model"
)

var (
	ErrNotFound      = errors.New("recipe not found")
	ErrEmptyLibrary  = errors.New("recipe library is empty")
)

type RecipeRepo interface {
	List(ctx context.Context) ([]*model.Recipe, error)
	GetBySlug(ctx context.Context, slug string) (*model.Recipe, error)
	Count(ctx context.Context) (int, error)
	BulkInsert(ctx context.Context, recipes []*model.Recipe) error
}
