package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
)

func AddProduct(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "add_product"

	case 1:
		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
