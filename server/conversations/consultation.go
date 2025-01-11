package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func RegisterConsultation(update tgbotapi.Update, ac app.App, state *app.UserState) {
	if !bot.IsSuperRole(update, ac) {
		return
	}

	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "register_consultation"

		bot.SendText(ac, update, "نام و نام خانوادگی خود را وارد کنید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		state.Data["name"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "لطفا شماره خود را وارد کنید")

		state.Step = 2
	case 2:
		r := regexp.MustCompile("^0[0-9]{2}[0-9]{8}$")

		if update.Message.Text == "" || !r.MatchString(update.Message.Text) {
			bot.SendText(ac, update, "لطفا یک شماره تلفن معتبر ارسال کنید!")
			return
		}

		state.Data["phone"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "لطفا پیام خود را برای مشاور بنویسید")
		state.Step = 3
	case 3:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		state.Data["message"] = strings.TrimSpace(update.Message.Text)

		// Send message
		//msg := tgbotapi.NewMessage(update.FromChat().ID, text)

		bot.SendText(ac, update, "پیام شما برای مشاور ارسال شد. در اولین فرصت با شما تماس گرفته خواهد شد.")

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
