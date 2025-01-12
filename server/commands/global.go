package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func Start(ac app.App, update tgbotapi.Update) {
	isSuper := bot.IsSuperRole(update, ac)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "سلام و درود به ربات طاها تهویه خوش آمدید. می‌توانید برای دریافت راهنمایی همین الان /help را وارد کنید.")

	msg.ReplyMarkup = keyboards.InlineKeyboard(menus.MainMenu, isSuper)

	bot.SendMessage(ac, msg)
}

func About(ac app.App, update tgbotapi.Update) {
	setting, err := ac.SettingsService().GetSetting("about")

	if err != nil {
		bot.SendText(ac, update, "خطایی در سرور رخ داد!")
		return
	}

	bot.SendMarkdown(ac, update, setting.Content.Content)
}

func Menu(ac app.App, update tgbotapi.Update) {
	isSuper := bot.IsSuperRole(update, ac)

	msg := tgbotapi.NewMessage(update.FromChat().ID, "منو خدمت شما")

	msg.ReplyMarkup = keyboards.InlineKeyboard(menus.MainMenu, isSuper)

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

func Support(ac app.App, update tgbotapi.Update) {
	bot.SendText(ac, update, "برای دریافت مشاوره، با آیدی @Taha_tahvieh در ارتباط باشید.")
}

func Help(ac app.App, update tgbotapi.Update) {
	help, err := ac.SettingsService().GetSetting("help")

	if err != nil {
		bot.SendText(ac, update, "خطایی در سرور رخ داد!")
		return
	}

	bot.SendMarkdown(ac, update, help.Content.Content)
}
