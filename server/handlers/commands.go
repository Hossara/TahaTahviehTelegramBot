package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/server/commands"
)

func HandleCommands(update tgbotapi.Update, ac app.App) {
	println(update.Message.From.UserName)

	cmdCfg := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "/start",
			Description: "شروع بات",
		},
		tgbotapi.BotCommand{
			Command:     "/menu",
			Description: "منو بات",
		},
		tgbotapi.BotCommand{
			Command:     "/product_list",
			Description: "لیست محصولات",
		},
		tgbotapi.BotCommand{
			Command:     "/register_consultation",
			Description: "ثبت‌نام برای مشاوره تلفنی",
		},
		tgbotapi.BotCommand{
			Command:     "/support",
			Description: "ارتباط‌ با پشتیبانی",
		},
		tgbotapi.BotCommand{
			Command:     "/faq",
			Description: "سوالات متداول",
		},
		tgbotapi.BotCommand{
			Command:     "/about",
			Description: "درباره ما",
		},
	)

	ac.Bot().Send(cmdCfg)

	switch update.Message.Command() {
	case "start":
		commands.CommandStart(ac, update)
	case "about":
		commands.CommandAbout(ac, update)
	}
	return
}
