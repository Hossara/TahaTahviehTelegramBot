package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/server/conversations"
)

func HandleConversations(update tgbotapi.Update, ac app.App) {
	userState := ac.AppState(update.SentFrom().ID)

	if userState.Active {
		switch userState.Conversation {
		case "update_about":
			conversations.UpdateAbout(update, ac, userState)
		}
	}
}
