package app

import (
	"context"
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/config"
	"taha_tahvieh_tg_bot/internal/settings"
	"taha_tahvieh_tg_bot/pkg/adapters/storage"
	"taha_tahvieh_tg_bot/pkg/postgres"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	settingsPort "taha_tahvieh_tg_bot/internal/settings/port"
)

type app struct {
	cfg             config.Config
	ctx             context.Context
	db              *gorm.DB
	bot             *tgbotapi.BotAPI
	settingsService settingsPort.Service
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Bot() *tgbotapi.BotAPI {
	return a.bot
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) SettingsService() settingsPort.Service {
	if a.settingsService == nil {
		if a.db != nil {
			a.settingsService = settings.NewService(a.ctx, storage.NewSettingRepo(a.db))

			if err := a.settingsService.RunMigrations(); err != nil {
				panic("failed to run migrations")
			}

			return a.settingsService
		}

		return nil
	}

	return a.settingsService
}

func (a *app) setBot() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Pass,
		Name:   a.cfg.DB.Name,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func NewApp(ctx context.Context, cfg config.Config, bot *tgbotapi.BotAPI) (App, error) {
	a := &app{cfg: cfg, bot: bot, ctx: ctx}

	if err := a.setBot(); err != nil {
		return nil, err
	}

	return a, nil
}

func MustNewApp(ctx context.Context, cfg config.Config, bot *tgbotapi.BotAPI) App {
	a, err := NewApp(ctx, cfg, bot)
	if err != nil {
		panic(err)
	}
	return a
}
