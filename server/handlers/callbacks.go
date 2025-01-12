package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
	"taha_tahvieh_tg_bot/server/conversations"
)

func HandleCallbacks(update tgbotapi.Update, ac app.App) {
	action := update.CallbackQuery.Data
	//chatID := update.CallbackQuery.Message.Chat.ID

	switch {
	case action == "/about":
		commands.CommandAbout(ac, update)

	case action == "/support":
		commands.Support(ac, update)

	case action == "/help":
		commands.Help(ac, update)

	case action == "/edit_about":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateAbout(update, ac, state)
	}
}
