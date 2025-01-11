package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func CommandStart(ac app.App, update tgbotapi.Update) {
	isSuper := bot.IsSuperRole(update, ac)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "سلام و درود به ربات طاها تهویه خوش آمدید. می‌توانید برای دریافت راهنمایی همین الان /help را وارد کنید.")

	msg.ReplyMarkup = keyboards.InlineKeyboard(menus.MainMenu, isSuper)

	bot.SendMessage(ac, msg)
}

func CommandAbout(ac app.App, update tgbotapi.Update) {
	setting, err := ac.SettingsService().GetSetting("about")

	if err != nil {
		bot.SendText(ac, update, "خطایی در سرور رخ داد!")
		return
	}

	bot.SendMarkdown(ac, update, setting.Content.Content)
}

func CommandMenu(ac app.App, update tgbotapi.Update) {
	isSuper := bot.IsSuperRole(update, ac)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "منو خدمت شما")

	msg.ReplyMarkup = keyboards.InlineKeyboard(menus.MainMenu, isSuper)

	bot.SendMessage(ac, msg)
}
