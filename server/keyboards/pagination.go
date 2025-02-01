package keyboards

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"taha_tahvieh_tg_bot/pkg/router"
	"taha_tahvieh_tg_bot/server/menus"
)

func getPaginationRow(currentPage, totalPages int, url, key string) []tgbotapi.InlineKeyboardButton {
	var paginationRow []tgbotapi.InlineKeyboardButton

	if totalPages > 1 {
		if currentPage > 1 {
			prev, err := router.ReplaceQueryParam(url, key, strconv.Itoa(currentPage-1))
			first, err := router.ReplaceQueryParam(url, key, strconv.Itoa(1))

			if err != nil {
				return nil
			}

			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData("<<", first)) // First
			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData("<", prev))   // Previous
		}

		paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d / %d", currentPage, totalPages), "page:current"))

		if currentPage < totalPages {
			next, err := router.ReplaceQueryParam(url, key, strconv.Itoa(currentPage+1))
			last, err := router.ReplaceQueryParam(url, key, strconv.Itoa(totalPages))

			if err != nil {
				return nil
			}

			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData(">", next))  // Next
			paginationRow = append(paginationRow, tgbotapi.NewInlineKeyboardButtonData(">>", last)) // Last
		}
	}

	return paginationRow
}

func InlinePaginationColumnKeyboard(
	menu []menus.MenuItem, isAdmin bool, currentPage,
	totalPages int, url, key string,
) tgbotapi.InlineKeyboardMarkup {
	var keyboard [][]tgbotapi.InlineKeyboardButton

	for _, item := range menu {
		if item.IsAdmin && !isAdmin {
			continue
		}

		keyboard = append(keyboard, []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData(item.Name, item.Path),
		})
	}

	pg := getPaginationRow(currentPage, totalPages, url, key)

	if len(pg) > 0 {
		keyboard = append(keyboard, pg)
	}

	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}

func InlinePaginationKeyboard(
	menu [][]menus.MenuItem, isAdmin bool, currentPage,
	totalPages int, url, key string,
) tgbotapi.InlineKeyboardMarkup {
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

	pg := getPaginationRow(currentPage, totalPages, url, key)

	if len(pg) > 0 {
		keyboardRows = append(keyboardRows, pg)
	}

	return tgbotapi.NewInlineKeyboardMarkup(keyboardRows...)
}
