package port

import (
	"taha_tahvieh_tg_bot/internal/faq/domain"
)

type Service interface {
	GetAllQuestions() ([]*domain.FrequentQuestion, error)
	GetQuestion(id domain.QuestionID) (*domain.FrequentQuestion, error)
	DeleteQuestion(id domain.QuestionID) error
	AddQuestion(question *domain.FrequentQuestion) error
	UpdateQuestion(question *domain.FrequentQuestion) error
	RunMigrations() error
}
