package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/settings/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func UpdateAbout(update tgbotapi.Update, ac app.App, state *app.UserState) {
	if !bot.IsSuperRole(update, ac) {
		return
	}

	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_about"

		bot.SendText(ac, update, "متن جدید را ارسال کنید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن معتبر ارسال کنید!")
			return
		}

		bot.SendText(ac, update, "در حال ویرایش متن درباره ما...")

		setting, err := ac.SettingsService().GetSetting("about")

		if err != nil {
			log.Println("Error while get about text", err)
			bot.SendText(ac, update, "خطا هنگام دریافت متن درباره ما")
			return
		}

		err = ac.SettingsService().UpdateSetting(&domain.Setting{
			SettingID: setting.SettingID,
			Title:     "about",
			Content: domain.Content{
				Content: strings.TrimSpace(update.Message.Text),
			},
		})

		if err != nil {
			log.Println("Error while updating about text", err)
			bot.SendText(ac, update, "خطا هنگام بروزرسانی متن درباره ما")
			return
		}

		bot.SendText(ac, update, "متن درباره ما با موفقیت بروزرسانی شد!")

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
