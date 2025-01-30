package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/faq/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/pkg/utils"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func GetFaqsMenu(questions []*domain.FrequentQuestion, action string) []menus.MenuItem {
	var menu []menus.MenuItem

	menu = utils.Map(questions, func(t *domain.FrequentQuestion) menus.MenuItem {
		return menus.MenuItem{
			Name:    t.Question,
			IsAdmin: false,
			Path:    fmt.Sprintf("/%s/%d", action, t.QuestionID),
		}
	})

	menu = append(menu, menus.MenuItem{
		Name:    "منو اصلی",
		IsAdmin: false,
		Path:    "/menu",
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

	msg.ReplyMarkup = keyboards.InlineKeyboardColumn(GetFaqsMenu(questions, "get_faq"), false)

	bot.SendMessage(ac, msg)
}

func QuestionRemoveFaq(ac app.App, update tgbotapi.Update, id string) {
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
		"آیا از انتخاب خود اطمینان دارید؟"+
			"سوال درحال حذف شدن:\n\n"+
			question.Question,
	)

	msg.ReplyMarkup = keyboards.InlineKeyboard([][]menus.MenuItem{
		{
			{
				Name:    "منو اصلی",
				Path:    "/menu",
				IsAdmin: false,
			},
			{
				Name:    "حذف کردن",
				Path:    fmt.Sprintf("/del_faq/%s", id),
				IsAdmin: false,
			},
		},
	}, false)

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

func FaqMenu(ac app.App, update tgbotapi.Update) {
	isSuper := bot.IsSuperRole(update, ac)

	if !isSuper {
		return
	}

	msg := tgbotapi.NewMessage(update.FromChat().ID, "منو سوالات پرتکرار خدمت شما")

	msg.ReplyMarkup = keyboards.InlineKeyboard(menus.FaqMenu, isSuper)

	bot.SendMessage(ac, msg)
}

func RemoveFaqMenu(ac app.App, update tgbotapi.Update) {
	questions, err := ac.FaqService().GetAllQuestions()

	if err != nil {
		bot.SendText(ac, update, "خطا هنگام دریافت سوالات!")
		return
	}

	msg := tgbotapi.NewMessage(update.FromChat().ID, "برای حذف هر سوال بر روی نام آن کلیک کنید")

	msg.ReplyMarkup = keyboards.InlineKeyboardColumn(GetFaqsMenu(questions, "q_r_faq"), false)

	bot.SendMessage(ac, msg)
}

func UpdateFaq(ac app.App, update tgbotapi.Update) {
	questions, err := ac.FaqService().GetAllQuestions()

	if err != nil {
		bot.SendText(ac, update, "خطا هنگام دریافت سوالات!")
		return
	}

	msg := tgbotapi.NewMessage(update.FromChat().ID, "برای ویرایش هر سوال بر روی عنوان آن کلیک کنید")

	msg.ReplyMarkup = keyboards.InlineKeyboardColumn(GetFaqsMenu(questions, "update_faq"), false)

	bot.SendMessage(ac, msg)
}

func RemoveFaq(ac app.App, update tgbotapi.Update, id string) {
	qId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		bot.SendText(ac, update, "سوال نامعتبر است!")
		return
	}

	bot.SendText(ac, update, "در حال حذف سوال...")

	err = ac.FaqService().DeleteQuestion(domain.QuestionID(qId))

	if err != nil {
		log.Printf("%v\n", err)
		bot.SendText(ac, update, "خطا هنگام حذف سوال")
		return
	}

	bot.SendText(ac, update, "سوال با موفقیت حذف شد!")
}
