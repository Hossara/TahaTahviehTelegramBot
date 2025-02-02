package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/pkg/utils"
	"taha_tahvieh_tg_bot/server/constants"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func ProductManagementMenu(ac app.App, update tgbotapi.Update, menu [][]menus.MenuItem) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, constants.MenuResponse)

	msg.ReplyMarkup = keyboards.InlineKeyboard(menu, true)

	bot.SendMessage(ac, msg)
}

func SearchProductMenu(ac app.App, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "دوست دارید بر اساس کدوم یک از فاکتور های زیر جستجو انجام دهید؟")

	msg.ReplyMarkup = keyboards.InlineKeyboard(menus.SearchProductMenu, true)

	bot.SendMessage(ac, msg)
}

func SelectProductMenu(ac app.App, update tgbotapi.Update, action, menu, text string, page, prev int, meta map[string]string) {
	send := func(keyboard tgbotapi.InlineKeyboardMarkup) {
		msg := tgbotapi.NewMessage(update.FromChat().ID, text)

		if page > 1 {
			bot.SendMessage(ac, tgbotapi.NewEditMessageReplyMarkup(update.FromChat().ID, prev, keyboard))
			return
		}

		msg.ReplyMarkup = keyboard
		bot.SendMessage(ac, msg)
	}

	addMain := func(menu []menus.MenuItem) []menus.MenuItem {
		return append(menu, menus.MenuItem{Name: "منو اصلی", IsAdmin: false, Path: "/menu"})
	}

	var menuItems []menus.MenuItem
	var keyboard tgbotapi.InlineKeyboardMarkup

	switch menu {
	case "brand":
		brands, err := ac.ProductService().GetAllBrands(page, 10)

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "خطا هنگام دریافت برند ها!")
			return
		}

		act := map[string]string{
			"search":      "/search/type",
			"add_product": "/product/product/add",
		}

		menuItems = utils.Map(brands.Data, func(t *productDomain.Brand) menus.MenuItem {
			return menus.MenuItem{
				Name: t.Title, IsAdmin: false,
				Path: fmt.Sprintf("%s?page=1&brand=%d", act[action], t.ID),
			}
		})

		keyboard = keyboards.InlinePaginationColumnKeyboard(
			addMain(menuItems), false,
			page, brands.Pages, fmt.Sprintf("/search/brand?page=%d", page), "page")
	case "type":
		brandID, err := strconv.Atoi(meta["brand"])

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "برند نامعتبر!")
			return
		}

		types, err := ac.ProductService().GetAllProductTypes(page, 10)

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "خطا هنگام دریافت دسته‌بندی ها!")
			return
		}

		act := map[string]string{
			"search":      "/search/product",
			"add_product": "/product/product/add",
		}

		menuItems = utils.Map(types.Data, func(t *productDomain.ProductType) menus.MenuItem {
			return menus.MenuItem{
				Name: t.Title, IsAdmin: false,
				Path: fmt.Sprintf("%s?page=1&brand=%d&type=%d", act[action], brandID, t.ID),
			}
		})

		keyboard = keyboards.InlinePaginationColumnKeyboard(
			addMain(menuItems), false,
			page, types.Pages, fmt.Sprintf("/search/type?page=%d&brand=%d", page, brandID), "page")
	}

	send(keyboard)
}
