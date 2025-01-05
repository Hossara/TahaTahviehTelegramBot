package port

import "taha_tahvieh_tg_bot/internal/settings/domain"

type Service interface {
	GetSetting(title string) (*domain.Setting, error)
	UpdateSetting(setting *domain.Setting) error
	RunMigrations() error
}
