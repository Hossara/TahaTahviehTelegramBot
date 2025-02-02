package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	router "taha_tahvieh_tg_bot/pkg/router"
	"taha_tahvieh_tg_bot/server/commands"
	"taha_tahvieh_tg_bot/server/conversations"
	"taha_tahvieh_tg_bot/server/menus"
)

func HandleCallbacks(update tgbotapi.Update, ac app.App) {
	action := update.CallbackQuery.Data
	r := router.NewRouter()

	// -------------------- General
	r.Handle("/about/{action}", func(vars router.PathVars, queries router.UrlQueries) {
		switch vars["action"] {
		case "get":
			commands.About(ac, update)
		case "update":
			if !bot.IsSuperRole(update, ac) {
				return
			}

			state := bot.ResetUserState(update, ac)
			conversations.UpdateAbout(update, ac, state)
		}
	})

	r.Handle("/menu", func(vars router.PathVars, queries router.UrlQueries) {
		commands.Menu(ac, update)
	})

	r.Handle("/support", func(vars router.PathVars, queries router.UrlQueries) {
		commands.Support(ac, update)
	})

	r.Handle("/help/{action}", func(vars router.PathVars, queries router.UrlQueries) {
		switch vars["action"] {
		case "get":
			commands.Help(ac, update)
		case "update":
			if !bot.IsSuperRole(update, ac) {
				return
			}

			state := bot.ResetUserState(update, ac)
			conversations.UpdateHelp(update, ac, state)
		}
	})

	// -------------------- Search
	r.Handle("/search", func(vars router.PathVars, queries router.UrlQueries) {
		commands.SearchProductMenu(ac, update)
	})

	r.Handle("/search/{query}", func(vars router.PathVars, queries router.UrlQueries) {
		pageQ, pageOk := queries["page"]

		page, pageErr := strconv.Atoi(pageQ)
		msID := update.CallbackQuery.Message.MessageID

		if !pageOk || pageErr != nil {
			bot.SendText(ac, update, "صفحه نامعتبر است.")
			return
		}

		switch vars["query"] {
		case "title":
		case "type":
			commands.SelectProductMenu(
				ac, update, "search", "type",
				"جهت دیدن محصولات هر دسته‌بندی بر روی عنوان آن کلیک کنید",
				page, msID, map[string]string{
					"brand": queries["brand"],
				},
			)
		case "brand":
			commands.SelectProductMenu(
				ac, update, "search", "brand",
				"برند مورد نظر خود را انتخاب کنید",
				page, msID, map[string]string{},
			)
		}
	})

	// -------------------- Product Management
	r.Handle("/manage/{query}", func(vars router.PathVars, queries router.UrlQueries) {
		if !bot.IsSuperRole(update, ac) {
			return
		}

		switch vars["query"] {
		case "brands":
			commands.ProductManagementMenu(ac, update, menus.ManageBrandMenu)
		case "product_types":
			commands.ProductManagementMenu(ac, update, menus.ManageProductTypeMenu)
		}
	})

	// -------------------- Products
	r.Handle("/product/{entity}/{action}", func(vars router.PathVars, queries router.UrlQueries) {
		if !bot.IsSuperRole(update, ac) {
			return
		}

		if vars["action"] == "add" {
			brandQ, brandOk := queries["brand"]
			typeQ, typeOk := queries["type"]

			brandID, brandErr := strconv.Atoi(brandQ)
			typeID, typeErr := strconv.Atoi(typeQ)

			var state *app.UserState

			if (brandOk && brandErr == nil) || (typeOk && typeErr == nil) {
				state = ac.AppState(update.SentFrom().ID)
			} else {
				state = bot.ResetUserState(update, ac)
			}

			switch vars["entity"] {
			case "product":
				conversations.AddProduct(update, ac, state, brandID, typeID)
			case "brand":
				conversations.AddBrand(update, ac, state)
			case "type":
				conversations.AddProductType(update, ac, state)
			}
			return
		}

		pageQ, pageOk := queries["page"]

		page, pageErr := strconv.Atoi(pageQ)

		if !pageOk || pageErr != nil {
			bot.SendText(ac, update, "صفحه نامعتبر است.")
			return
		}

		actionText := map[string]string{
			"remove": "حذف", "update": "ویرایش",
		}

		categoryText := map[string]string{
			"brand": "برند", "type": "دسته‌بندی محصول",
		}

		commands.SelectProductMenu(
			ac, update, vars["action"], vars["entity"],
			fmt.Sprintf(
				"جهت %s هر %s، بر روی نام آن کلیک کنید.",
				actionText[vars["action"]], categoryText[vars["entity"]],
			),
			page, update.CallbackQuery.Message.MessageID,
			map[string]string{},
		)
	})

	// -------------------- FAQ
	r.Handle("/faq", func(vars router.PathVars, queries router.UrlQueries) {
		commands.FaqList(ac, update)
	})

	r.Handle("/faq/add", func(vars router.PathVars, queries router.UrlQueries) {
		if !bot.IsSuperRole(update, ac) {
			return
		}

		state := bot.ResetUserState(update, ac)
		conversations.AddFaq(update, ac, state)
	})

	r.Handle("/faq/menu/{action}", func(vars router.PathVars, queries router.UrlQueries) {
		if !bot.IsSuperRole(update, ac) {
			return
		}

		switch vars["action"] {
		case "":
			commands.FaqMenu(ac, update)
		case "update":
			commands.UpdateFaqMenu(ac, update)
		case "remove":
			commands.RemoveFaqMenu(ac, update)
		}
	})

	r.Handle("/faq/{action}/{question}", func(vars router.PathVars, queries router.UrlQueries) {
		id, err := strconv.ParseUint(vars["question"], 10, 64)

		if err != nil {
			bot.SendText(ac, update, "سوال نامعتبر است!")
			return
		}

		switch vars["action"] {
		case "get":
			commands.GetFaq(ac, update, id)
		case "update":
			if !bot.IsSuperRole(update, ac) {
				return
			}
			state := bot.ResetUserState(update, ac)
			state.Data["id"] = strconv.FormatUint(id, 10)
			conversations.UpdateFaq(update, ac, state)
		case "remove":
			if !bot.IsSuperRole(update, ac) {
				return
			}
			commands.RemoveFaq(ac, update, id)
		case "remove_confirm":
			if !bot.IsSuperRole(update, ac) {
				return
			}
			commands.QuestionRemoveFaq(ac, update, id)
		}
	})

	// Parse & Call
	r.Parse(action)
}
