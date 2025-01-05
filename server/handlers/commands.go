package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
)

func handleMenu(update tgbotapi.Update, ac app.App) {
	botCommands := commands.BotCommands

	// Commands for admins & super admins
	if bot.IsSuperRole(update, ac) {
		botCommands = append(botCommands, commands.AdminCommands...)
	}

	cmdCfg := tgbotapi.NewSetMyCommands(botCommands...)

	bot.SendRequest(ac, cmdCfg)
}

func HandleCommands(update tgbotapi.Update, ac app.App) {
	// Handle command menu
	handleMenu(update, ac)

	switch update.Message.Command() {
	case "start":
		commands.CommandStart(ac, update)
	case "about":
		commands.CommandAbout(ac, update)
	}
	return
}
