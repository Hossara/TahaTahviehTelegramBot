package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/config"
)

type app struct {
	cfg config.Config
	bot *tgbotapi.BotAPI
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Bot() *tgbotapi.BotAPI {
	return a.bot
}

func (a *app) setBot(bot *tgbotapi.BotAPI) {
	a.bot = bot
}

func NewApp(cfg config.Config, bot *tgbotapi.BotAPI) (App, error) {
	a := &app{cfg: cfg}

	a.setBot(bot)

	/*	if err := a.setDB(); err != nil {
			return nil, err
		}

		a.setRedis()
		a.setEmailService()*/

	return a, nil
}

func MustNewApp(cfg config.Config, bot *tgbotapi.BotAPI) App {
	a, err := NewApp(cfg, bot)
	if err != nil {
		panic(err)
	}
	return a
}
