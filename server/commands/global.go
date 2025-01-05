package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func CommandStart(ac app.App, update tgbotapi.Update) {

	/*	isAdmin := super_admin.IsSuperAdmin(update.Message.From.ID)

		if isAdmin {
			err := users.UpdateChatID(database.DB, update.SentFrom().ID, update.Message.Chat.ID)
			if err != nil {
				botLogger.Error(
					"Error while updating chatID for User",
					zap.Error(err),
					zap.String("user_id", strconv.FormatInt(update.SentFrom().ID, 10)),
				)
			}
		}
		_, err := users.LoginUser(database.DB, update.SentFrom().ID, update.Message.Chat.ID)

		if err != nil {
			botLogger.Error(
				"Error while calling login user",
				zap.Error(err),
				zap.String("user_id", strconv.FormatInt(update.SentFrom().ID, 10)),
			)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "خطا نگام دریافت اطلاعات کاربر")
			bot.Send(msg)
			return
		}*/

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "سلام و درود به ربات طاها تهویه خوش آمدید. می‌توانید برای دریافت راهنمایی همین الان /help را وارد کنید.")

	//msg.ReplyMarkup = keyboards.InlineKeyboard(menus.MainMenu, isAdmin)

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
