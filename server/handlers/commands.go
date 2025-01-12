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
	// -------------------- General
	case "start":
		commands.Start(ac, update)
	case "about":
		commands.About(ac, update)
	case "menu":
		commands.Menu(ac, update)
	case "support":
		commands.Support(ac, update)
	case "help":
		commands.Help(ac, update)

	// -------------------- FAQ
	case "faq":
		commands.Faq(ac, update)

	case "faq_menu":
		commands.FaqMenu(ac, update)

	case "add_faq":
		state := bot.ResetUserState(update, ac)
		conversations.AddFaq(update, ac, state)

	case "remove_faq":
		state := bot.ResetUserState(update, ac)
		conversations.RemoveFaq(update, ac, state)

	case "update_faq":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateFaq(update, ac, state)

	// -------------------- General Conversations
	case "edit_about":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateAbout(update, ac, state)

	case "edit_help":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateHelp(update, ac, state)
	}
	return
}
