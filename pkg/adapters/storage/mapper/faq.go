package mapper

import (
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/models"
	"time"
)

func ToDomainFaq(m *models.Faq) *domain.FrequentQuestion {
	if m == nil {
		return nil
	}

	return &domain.FrequentQuestion{
		QuestionID: domain.QuestionID(m.ID),
		Question:   m.Question,
		Answer:     m.Answer,
	}
}

func ToModelFaq(m *domain.FrequentQuestion) *models.Faq {
	if m == nil {
		return nil
	}

	return &models.Faq{
		ID:        uint8(m.QuestionID),
		Question:  m.Question,
		Answer:    m.Answer,
		UpdatedAt: time.Now(),
	}
}
