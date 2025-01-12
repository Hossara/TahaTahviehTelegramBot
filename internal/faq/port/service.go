package port

import (
	"context"
	"taha_tahvieh_tg_bot/internal/faq/domain"
)

type Repo interface {
	FindAll(ctx context.Context) ([]*domain.FrequentQuestion, error)
	Insert(ctx context.Context, question *domain.FrequentQuestion) (domain.QuestionID, error)
	FindByID(ctx context.Context, id domain.QuestionID) (*domain.FrequentQuestion, error)
	RemoveByID(ctx context.Context, id domain.QuestionID) error
	Update(ctx context.Context, question *domain.FrequentQuestion) error
	RunMigrations() error
}
