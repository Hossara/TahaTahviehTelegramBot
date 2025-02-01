package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func AddFaq(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "add_faq"

		bot.SendText(ac, update, "عنوان سوال را بنویسید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		state.Data["question"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "حالا متن پاسخ را بنویسید")
		state.Step = 2
	case 2:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		bot.SendText(ac, update, "در حال افزودن سوال...")

		err := ac.FaqService().AddQuestion(&domain.FrequentQuestion{
			Question: state.Data["question"],
			Answer:   strings.TrimSpace(update.Message.Text),
		})

		if err != nil {
			log.Println("Error while insert new faq", err)
			bot.SendText(ac, update, "خطا هنگام افزودن سوال جدید!")
			return
		}

		bot.SendText(ac, update, "سوال با موفقیت اضافه شد.")

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
