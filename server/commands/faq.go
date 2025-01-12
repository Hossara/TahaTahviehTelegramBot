package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func getFaqsMenu(questions []*domain.FrequentQuestion) [][]menus.MenuItem {
	var menu [][]menus.MenuItem

	for i := 0; i < len(questions); i += 2 {
		var row []menus.MenuItem

		row = append(row, menus.MenuItem{
			Name:    questions[i].Question,
			IsAdmin: false,
			Path:    fmt.Sprintf("/get_faq/%d", questions[i].QuestionID),
		})

		if i+1 < len(questions) {
			row = append(row, menus.MenuItem{
				Name:    questions[i+1].Question,
				IsAdmin: false,
				Path:    fmt.Sprintf("/get_faq/%d", questions[i+1].QuestionID),
			})
		}

		menu = append(menu, row)
	}

	menu = append(menu, []menus.MenuItem{
		{
			Name:    "منو اصلی",
			IsAdmin: false,
			Path:    "/menu",
		},
	})

	return menu
}

func FaqList(ac app.App, update tgbotapi.Update) {
	questions, err := ac.FaqService().GetAllQuestions()

	if err != nil {
		bot.SendText(ac, update, "خطا هنگام دریافت سوالات!")
		return
	}

	msg := tgbotapi.NewMessage(update.FromChat().ID, "برای دیدن پاسخ هر سوال بر روی عنوان آن کلیک کنید")

	msg.ReplyMarkup = keyboards.InlineKeyboard(getFaqsMenu(questions), false)

	bot.SendMessage(ac, msg)
}

func Faq(ac app.App, update tgbotapi.Update, id string) {
	qId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		bot.SendText(ac, update, "سوال نامعتبر است!")
		return
	}

	question, err := ac.FaqService().GetQuestion(domain.QuestionID(qId))

	if err != nil {
		bot.SendText(ac, update, "خطا هنگام دریافت اطلاعات سوال. سوال یافت نشد.")
		return
	}

	msg := tgbotapi.NewMessage(
		update.FromChat().ID,
		question.Question+"\n\n"+"پاسخ: \n"+question.Answer,
	)

	msg.ReplyMarkup = keyboards.InlineKeyboard([][]menus.MenuItem{
		{
			menus.MenuItem{
				Name:    "منو اصلی",
				IsAdmin: false,
				Path:    "/menu",
			},
			menus.MenuItem{
				Name:    "منو سوالات",
				IsAdmin: false,
				Path:    "/faq",
			},
		},
	}, false)

	msg.ParseMode = tgbotapi.ModeMarkdown

	bot.SendMessage(ac, msg)
}
