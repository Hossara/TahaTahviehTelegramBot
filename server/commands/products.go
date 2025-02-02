package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	psDomain "taha_tahvieh_tg_bot/internal/product_storage/domain"
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

func send(ac app.App, update tgbotapi.Update, text string, page, prev int, keyboard tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, text)

	if page > 1 {
		bot.SendMessage(ac, tgbotapi.NewEditMessageReplyMarkup(update.FromChat().ID, prev, keyboard))
		return
	}

	msg.ReplyMarkup = keyboard
	bot.SendMessage(ac, msg)
}

func addMain(menu []menus.MenuItem) []menus.MenuItem {
	return append(menu, menus.MenuItem{Name: "منو اصلی", IsAdmin: false, Path: "/menu"})
}

func GetProduct(ac app.App, update tgbotapi.Update, id int64) {
	product, err := ac.ProductService().GetProduct(productDomain.ProductID(id))

	if err != nil {
		log.Println(err)
		bot.SendText(ac, update, "خطا هنگام خواندن اطلاعات محصول!")
		return
	}

	msg := tgbotapi.NewMessage(update.FromChat().ID, fmt.Sprintf(
		"نام محصول: %s\n"+
			"کد محصول: %s\n"+
			"برند: %s\n"+
			"دسته‌بندی: %s\n"+
			"توضیحات محصول: \n%s\n",
		product.Title,
		product.UUID,
		product.Brand.Title,
		product.Type.Title,
		product.Description,
	))

	msg.ReplyMarkup = keyboards.InlineKeyboardColumn([]menus.MenuItem{
		{Path: "/product/", Name: "ویرایش محصول", IsAdmin: false},
		{Path: "/product/", Name: "دریافت فایل های محصول", IsAdmin: false},
		{Path: "/product/", Name: "حذف محصول", IsAdmin: false},
	}, false)

	bot.SendMessage(ac, msg)
}

func ProductList(ac app.App, update tgbotapi.Update, brandID int64, typeID int64, title string, page, prev int) {
	var menuItems []menus.MenuItem
	var keyboard tgbotapi.InlineKeyboardMarkup

	product, err := ac.ProductService().GetAllProductsBasedOn(
		productDomain.BrandID(brandID),
		productDomain.ProductTypeID(typeID),
		title, page, 10,
	)

	if err != nil {
		log.Println(err)
		bot.SendText(ac, update, "خطا هنگام دریافت محصولات!")
		return
	}

	menuItems = utils.Map(product.Data, func(t *psDomain.Product) menus.MenuItem {
		return menus.MenuItem{
			Name: t.Title, IsAdmin: false,
			Path: fmt.Sprintf("/product/product/get?pid=%d", t.ID),
		}
	})

	keyboard = keyboards.InlinePaginationColumnKeyboard(
		addMain(menuItems), false,
		page, product.Pages, fmt.Sprintf("/search/product?page=%d", page), "page")
	send(ac, update, "جهت مشاهده هر محصول بر روی عنوان آن کلیک کنید.", page, prev, keyboard)

}

func SelectProductMenu(ac app.App, update tgbotapi.Update, action, menu, text string, page, prev int, meta map[string]string) {
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

	send(ac, update, text, page, prev, keyboard)
}
