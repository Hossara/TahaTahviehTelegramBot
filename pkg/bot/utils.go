package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"taha_tahvieh_tg_bot/app"
)

func sendMessage(ac app.App, update tgbotapi.Update, text string, html bool) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	if html {
		msg.ParseMode = tgbotapi.ModeMarkdown
	}

	_, err := ac.Bot().Send(msg)

	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
}

func SendText(ac app.App, update tgbotapi.Update, text string) {
	sendMessage(ac, update, text, false)
}

func SendHtml(ac app.App, update tgbotapi.Update, text string) {
	sendMessage(ac, update, text, true)
}
