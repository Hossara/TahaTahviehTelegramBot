package migrations

import (
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/helpers"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/models"

	"github.com/go-gormigrate/gormigrate/v2"
)

func GetFaqMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: helpers.GenerateMigrationID("faq_table"),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Setting{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("faq")
			},
		},
	}
}
