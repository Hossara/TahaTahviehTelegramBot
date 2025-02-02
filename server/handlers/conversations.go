package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/server/conversations"
)

var gaps = map[string]func(update tgbotapi.Update, ac app.App, state *app.UserState){
	"add_brand":           conversations.AddBrand,
	"add_product_type":    conversations.AddProductType,
	"update_product_type": conversations.UpdateProductType,
	"update_brand":        conversations.UpdateBrand,
	"remove_product":      conversations.RemoveProduct,
	"update_about":        conversations.UpdateAbout,
	"update_help":         conversations.UpdateHelp,
	"add_faq":             conversations.AddFaq,
	"update_faq":          conversations.UpdateFaq,
}

func HandleConversations(update tgbotapi.Update, ac app.App) {
	userState := ac.AppState(update.SentFrom().ID)

	if userState.Active {
		fun, ok := gaps[userState.Conversation]

		if ok {
			fun(update, ac, userState)
		}

		switch userState.Conversation {
		case "add_product":
			conversations.AddProduct(update, ac, userState, 0, 0, 1, 0)
		}
	}
}
