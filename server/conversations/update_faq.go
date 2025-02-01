package conversations

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func UpdateFaq(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_faq"

		qId, err := strconv.ParseUint(state.Data["id"], 10, 64)

		if err != nil {
			bot.SendText(ac, update, "سوال نامعتبر است!")
			return
		}

		question, err := ac.FaqService().GetQuestion(domain.QuestionID(qId))

		if err != nil {
			bot.SendText(ac, update, "خطا هنگام دریافت اطلاعات سوال. سوال یافت نشد.")
			return
		}

		bot.SendMarkdown(
			ac, update,
			"عنوان جدید سوال را بنویسید:\n\n"+
				fmt.Sprintf("عنوان قبلی: %s", question.Question)+"\n\n"+
				"**نکته: برای بدون تغییر ماندن عنوان، عدد 1 انگلیسی را ارسال کنید!**",
		)

		state.Data["question"] = question.Question
		state.Data["answer"] = question.Answer
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		if update.Message.Text != "1" {
			state.Data["question"] = strings.TrimSpace(update.Message.Text)
		} else {
			bot.SendText(ac, update, "عنوان سوال بدون تغییر ماند.")
		}

		bot.SendMarkdown(
			ac, update,
			"حالا پاسخ جدید سوال را بنویسید:\n\n"+
				fmt.Sprintf("پاسخ قبلی: %s", state.Data["answer"])+"\n\n"+
				"**نکته: برای بدون تغییر ماندن عنوان، عدد 1 انگلیسی را ارسال کنید!**",
		)

		state.Step = 2

	case 2:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		if update.Message.Text != "1" {
			bot.SendText(ac, update, "درحال ویرایش سوال...")

			qId, err := strconv.ParseUint(state.Data["id"], 10, 64)

			if err != nil {
				bot.SendText(ac, update, "سوال نامعتبر است!")
				return
			}

			err = ac.FaqService().UpdateQuestion(&domain.FrequentQuestion{
				Question:   state.Data["question"],
				Answer:     update.Message.Text,
				QuestionID: domain.QuestionID(qId),
			})

			if err != nil {
				log.Println(err)
				bot.SendText(ac, update, "خطا هنگام بروزرسانی سوال")
				return
			}
		} else {
			bot.SendText(ac, update, "پاسخ سوال بدون تغییر ماند.")
		}

		bot.SendText(ac, update, "سوال با موفقیت ویرایش شد!")

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
