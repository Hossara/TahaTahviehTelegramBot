package port

import (
	"context"
	"taha_tahvieh_tg_bot/internal/settings/domain"
)

type Repo interface {
	FindSettingByTitle(ctx context.Context, title string) (*domain.Setting, error)
	UpdateSetting(ctx context.Context, setting domain.Setting) error
	RunMigrations() error
}
