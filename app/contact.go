package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"taha_tahvieh_tg_bot/config"
	settingsPort "taha_tahvieh_tg_bot/internal/settings/port"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Bot() *tgbotapi.BotAPI
	SettingsService() settingsPort.Service
}
