package settings

import (
	"context"
	"taha_tahvieh_tg_bot/internal/settings/domain"
	"taha_tahvieh_tg_bot/internal/settings/port"
)

type service struct {
	repo port.Repo
	ctx  context.Context
}

func NewService(ctx context.Context, repo port.Repo) port.Service {
	return &service{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *service) RunMigrations() error {
	return s.repo.RunMigrations()
}

func (s *service) GetSetting(title string) (*domain.Setting, error) {
	setting, err := s.repo.FindByTitle(s.ctx, title)

	if err != nil {
		return nil, err
	}

	return setting, nil
}

func (s *service) UpdateSetting(setting *domain.Setting) error {
	err := s.repo.Update(s.ctx, setting)

	if err != nil {
		return err
	}

	return nil
}
