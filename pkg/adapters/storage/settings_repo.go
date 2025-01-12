package storage

import (
	"context"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/internal/settings/domain"
	"taha_tahvieh_tg_bot/internal/settings/port"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/mapper"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/migrations"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/models"
)

type settingRepo struct {
	db *gorm.DB
}

func NewSettingRepo(db *gorm.DB) port.Repo {
	return &settingRepo{db}
}

func (r *settingRepo) FindByTitle(ctx context.Context, title string) (*domain.Setting, error) {
	var setting models.Setting

	err := r.db.WithContext(ctx).Where("title = ?", title).First(&setting).Error

	return mapper.ToDomainSetting(&setting), err
}

func (r *settingRepo) Update(ctx context.Context, setting *domain.Setting) error {
	modelSetting := mapper.ToModelSetting(setting)

	err := r.db.WithContext(ctx).Save(modelSetting).Error

	return err
}

func (r *settingRepo) RunMigrations() error {
	migrator := gormigrate.New(r.db, gormigrate.DefaultOptions, migrations.GetSettingMigrations())
	return migrator.Migrate()
}
