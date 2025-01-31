package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/constants"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func ProductManagementMenu(ac app.App, update tgbotapi.Update, menu [][]menus.MenuItem) {
	isSuper := bot.IsSuperRole(update, ac)

	if !isSuper {
		return
	}

	msg := tgbotapi.NewMessage(update.FromChat().ID, constants.MenuResponse)

	msg.ReplyMarkup = keyboards.InlineKeyboard(menu, true)

	bot.SendMessage(ac, msg)
}

func SearchProductMenu(ac app.App, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "دوست دارید بر اساس کدوم یک از فاکتور های زیر جستجو انجام دهید؟")

	msg.ReplyMarkup = keyboards.InlineKeyboard(menus.SearchProductMenu, true)

	bot.SendMessage(ac, msg)
}

func SelectProductMenu(ac app.App, update tgbotapi.Update, menu, title string) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, fmt.Sprintf("جهت دیدن %s بر روی عنوان آن کلیک کنید", title))

	/*var menuItems []menus.MenuItem
	//products, err := ac.ProductService()

	switch menu {
	case "brand":
		brands, err := ac.ProductService().GetAllBrands()

		menuItems = utils.Map(items, func(t *T) menus.MenuItem {
			return menus.MenuItem{
				Name:    t.Title,
				IsAdmin: false,
				Path:    fmt.Sprintf("/%s/%d", action, t.ID),
			}
		})
	case "type":

	}*/

	/*

		menu = append(menu, menus.MenuItem{Name: "منو اصلی", IsAdmin: false, Path: "/menu"})*/

	//msg.ReplyMarkup = keyboards.InlineKeyboardColumn(GetProductMenu(questions, action), false)

	bot.SendMessage(ac, msg)
}
