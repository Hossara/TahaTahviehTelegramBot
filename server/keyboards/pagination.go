package keyboards

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/server/menus"
)

func getPaginationRow(currentPage, totalPages int) []tgbotapi.InlineKeyboardButton {
	var paginationRow []tgbotapi.InlineKeyboardButton

	if totalPages > 1 {
		if currentPage > 1 {
			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData("<<", "page:1"))                             // First
			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData("<", fmt.Sprintf("page:%d", currentPage-1))) // Previous
		}

		paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d / %d", currentPage, totalPages), "page:current"))

		if currentPage < totalPages {
			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData(">", fmt.Sprintf("page:%d", currentPage+1))) // Next
			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData(">>", fmt.Sprintf("page:%d", totalPages)))   // Last
		}
	}

	return paginationRow
}

func InlinePaginationColumnKeyboard(menu []menus.MenuItem, isAdmin bool, currentPage, totalPages int) tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for _, item := range menu {
		if item.IsAdmin && !isAdmin {
			continue
		}

		keyboard = append(keyboard, []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData(item.Name, item.Path),
		})
	}

	keyboard = append(keyboard, getPaginationRow(currentPage, totalPages))

	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}

func InlinePaginationKeyboard(menu [][]menus.MenuItem, isAdmin bool, currentPage, totalPages int) tgbotapi.InlineKeyboardMarkup {
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

	keyboardRows = append(keyboardRows, getPaginationRow(currentPage, totalPages))

	// Return the complete keyboard markup
	return tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
}
