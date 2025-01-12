package faq

import (
	"context"
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/internal/faq/port"
)

func NewService(ctx context.Context, repo port.Repo) port.Service {
	return &service{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *service) RunMigrations() error {
	return s.repo.RunMigrations()
}

type service struct {
	repo port.Repo
	ctx  context.Context
}

func (s *service) AddQuestion(question *domain.FrequentQuestion) error {
	_, err := s.repo.Insert(s.ctx, question)

	return err
}

func (s *service) GetAllQuestions() ([]*domain.FrequentQuestion, error) {
	return s.repo.FindAll(s.ctx)
}

func (s *service) GetQuestion(id domain.QuestionID) (*domain.FrequentQuestion, error) {
	return s.repo.FindByID(s.ctx, id)
}
func (s *service) DeleteQuestion(id domain.QuestionID) error {
	return s.repo.RemoveByID(s.ctx, id)
}

func (s *service) UpdateQuestion(question *domain.FrequentQuestion) error {
	return s.repo.Update(s.ctx, question)
}
