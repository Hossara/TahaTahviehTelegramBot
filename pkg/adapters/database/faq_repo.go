package database

import (
	"context"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/internal/faq/port"
	"taha_tahvieh_tg_bot/pkg/adapters/database/helpers"
	"taha_tahvieh_tg_bot/pkg/adapters/database/mapper"
	"taha_tahvieh_tg_bot/pkg/adapters/database/models"
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

func (r *faqRepo) FindByID(ctx context.Context, id domain.QuestionID) (*domain.FrequentQuestion, error) {
	var question models.Faq

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&question).Error

	return mapper.ToDomainFaq(&question), err
}

func (r *faqRepo) RemoveByID(ctx context.Context, id domain.QuestionID) error {
	err := r.db.WithContext(ctx).Delete(&models.Faq{}, uint64(id)).Error

	return err
}

func (r *faqRepo) Update(ctx context.Context, question *domain.FrequentQuestion) error {
	err := r.db.WithContext(ctx).Save(mapper.ToModelFaq(question)).Error

	return err
}

func (r *faqRepo) RunMigrations() error {
	migrator := gormigrate.New(
		r.db, gormigrate.DefaultOptions,
		helpers.GetMigrations[models.Faq]("faq", &models.Faq{}),
	)

	return migrator.Migrate()
}
