package storage

import (
	"context"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/internal/faq/port"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/mapper"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/migrations"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/models"
	"taha_tahvieh_tg_bot/pkg/utils"
)

func NewFaqRepo(db *gorm.DB) port.Repo {
	return &faqRepo{db}
}

type faqRepo struct {
	db *gorm.DB
}

func (r *faqRepo) Insert(ctx context.Context, question *domain.FrequentQuestion) (domain.QuestionID, error) {
	q := mapper.ToModelFaq(question)

	return domain.QuestionID(q.ID), r.db.WithContext(ctx).Create(q).Error
}

func (r *faqRepo) FindAll(ctx context.Context) ([]*domain.FrequentQuestion, error) {
	var faqs []models.Faq

	err := r.db.WithContext(ctx).Find(&faqs).Error

	return utils.Map(faqs, func(t models.Faq) *domain.FrequentQuestion {
		return mapper.ToDomainFaq(&t)
	}), err
}

func (r *faqRepo) FindByTitle(ctx context.Context, title string) (*domain.FrequentQuestion, error) {
	//TODO implement me
	panic("implement me")
}

func (r *faqRepo) Update(ctx context.Context, question *domain.FrequentQuestion) error {
	//TODO implement me
	panic("implement me")
}

func (r *faqRepo) RunMigrations() error {
	migrator := gormigrate.New(r.db, gormigrate.DefaultOptions, migrations.GetFaqMigrations())
	return migrator.Migrate()
}
