package port

import (
	"context"
	"taha_tahvieh_tg_bot/internal/faq/domain"
)

type Repo interface {
	FindAll(ctx context.Context) ([]*domain.FrequentQuestion, error)
	Insert(ctx context.Context, question *domain.FrequentQuestion) (domain.QuestionID, error)
	FindByTitle(ctx context.Context, title string) (*domain.FrequentQuestion, error)
	Update(ctx context.Context, question *domain.FrequentQuestion) error
	RunMigrations() error
}
