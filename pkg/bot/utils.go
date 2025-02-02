package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"slices"
	"taha_tahvieh_tg_bot/app"
)

func sendMessage(ac app.App, update tgbotapi.Update, text string, markdown bool) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, text)

	if markdown {
		msg.ParseMode = tgbotapi.ModeMarkdown
	}

	_, err := ac.Bot().Send(msg)

	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
}

func SendMessage(ac app.App, c tgbotapi.Chattable) {
	_, err := ac.Bot().Send(c)

	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
}
func SendMessageReturns(ac app.App, c tgbotapi.Chattable) int {
	id, err := ac.Bot().Send(c)

	if err != nil {
		log.Println("Error sending message:", err)
		return 0
	}

	return id.MessageID
}

func SendRequest(ac app.App, c tgbotapi.Chattable) {
	_, err := ac.Bot().Request(c)

	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
}

func SendText(ac app.App, update tgbotapi.Update, text string) {
	sendMessage(ac, update, text, false)
}

func SendMarkdown(ac app.App, update tgbotapi.Update, text string) {
	sendMessage(ac, update, text, true)
}

func IsAdmin(update tgbotapi.Update, ac app.App) bool {
	constants := ac.Config().Constants
	username := update.SentFrom().UserName

	return slices.Contains(constants.Admins, username)
}

func IsSuperAdmin(update tgbotapi.Update, ac app.App) bool {
	constants := ac.Config().Constants
	username := update.SentFrom().UserName

	return slices.Contains(constants.SuperAdmins, username)
}

func IsSuperRole(update tgbotapi.Update, ac app.App) bool {
	return IsAdmin(update, ac) || IsSuperAdmin(update, ac)
}

func ResetUserState(update tgbotapi.Update, ac app.App) *app.UserState {
	userId := update.SentFrom().ID

	ac.ResetUserState(userId)
	return ac.AppState(userId)
}
