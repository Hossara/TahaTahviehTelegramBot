package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func RemoveFaq(update tgbotapi.Update, ac app.App, state *app.UserState) {
	if !bot.IsSuperRole(update, ac) {
		return
	}

	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "remove_faq"

		state.Step = 1
	case 1:

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
