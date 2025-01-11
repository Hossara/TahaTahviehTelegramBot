package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/server/menus"
)

func InlineKeyboard(menu [][]menus.MenuItem, isAdmin bool) tgbotapi.InlineKeyboardMarkup {
	var keyboardRows [][]tgbotapi.InlineKeyboardButton

	for _, items := range menu {
		var row []tgbotapi.InlineKeyboardButton

		for _, item := range items {

			if item.IsAdmin && !isAdmin {
				continue
			}
			row = append(row, tgbotapi.NewInlineKeyboardButtonData(item.Name, item.Path))
		}

		if len(row) > 0 {
			keyboardRows = append(keyboardRows, row)
		}
	}

	// Return the complete keyboard markup
	return tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
}
