package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/config"
)

type App interface {
	Config() config.Config
	Bot() *tgbotapi.BotAPI
}
