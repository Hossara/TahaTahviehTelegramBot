package port

import (
	"context"
	"taha_tahvieh_tg_bot/internal/settings/domain"
)

type Repo interface {
	FindByTitle(ctx context.Context, title string) (*domain.Setting, error)
	Update(ctx context.Context, setting *domain.Setting) error
	RunMigrations() error
}
