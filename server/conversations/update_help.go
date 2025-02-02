package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/settings/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func UpdateHelp(update tgbotapi.Update, ac app.App, state *app.UserState) {
	if !bot.IsSuperRole(update, ac) {
		return
	}

	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_help"

		bot.SendText(ac, update, "متن جدید را ارسال کنید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		bot.SendText(ac, update, "در حال ویرایش متن راهنمای ربات...")

		setting, err := ac.SettingsService().GetSetting("help")

		if err != nil {
			log.Println("Error while get about text", err)
			bot.SendText(ac, update, "خطا هنگام دریافت متن راهنمایی")
			return
		}

		err = ac.SettingsService().UpdateSetting(&domain.Setting{
			SettingID: setting.SettingID,
			Title:     "help",
			Content: domain.Content{
				Content: strings.TrimSpace(update.Message.Text),
			},
		})

		if err != nil {
			log.Println("Error while updating about text", err)
			bot.SendText(ac, update, "خطا هنگام بروزرسانی متن راهنما")
			return
		}

		bot.SendText(ac, update, "متن راهنمای ربات با موفقیت بروزرسانی شد!")

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
