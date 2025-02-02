package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
)

func SearchByTitle(update tgbotapi.Update, ac app.App, state *app.UserState, page, prev int) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "search_by_title"

		bot.SendText(ac, update, "نام محصول را بنویسید")
		state.Step = 1

	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		text := strings.TrimSpace(update.Message.Text)
		state.Data["text"] = text

		commands.ProductList(ac, update, 0, 0, text, page, prev)
	}
}
