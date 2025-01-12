package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
	"taha_tahvieh_tg_bot/server/conversations"
)

func HandleCommands(update tgbotapi.Update, ac app.App) {
	switch update.Message.Command() {
	case "start":
		commands.CommandStart(ac, update)
	case "about":
		commands.CommandAbout(ac, update)
	case "menu":
		commands.CommandMenu(ac, update)
	case "support":
		commands.Support(ac, update)
	case "help":
		commands.Help(ac, update)

	case "edit_about":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateAbout(update, ac, state)

	case "edit_help":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateHelp(update, ac, state)
	}
	return
}
