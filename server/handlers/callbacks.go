package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/pkg/router"
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
			state := bot.ResetUserState(update, ac)
			conversations.SearchByTitle(update, ac, state, page, msID)
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

		case "product":
			brandQ, brandOk := queries["brand"]
			typeQ, typeOk := queries["type"]
			title, titleOk := queries["title"]

			brandID, brandErr := strconv.Atoi(brandQ)
			typeID, typeErr := strconv.Atoi(typeQ)

			if titleOk {
				commands.ProductList(ac, update, 0, 0, title, page, msID)
			} else if typeOk && typeErr == nil && brandOk && brandErr == nil {
				commands.ProductList(ac, update, int64(brandID), int64(typeID), "", page, msID)
			} else {
				log.Println(typeErr)
				log.Println(brandErr)
			}
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

		pageQ, pageOk := queries["page"]
		page, pageErr := strconv.Atoi(pageQ)
		msID := update.CallbackQuery.Message.MessageID

		if vars["action"] == "add" {
			brandQ, brandOk := queries["brand"]
			typeQ, typeOk := queries["type"]

			brandID, brandErr := strconv.Atoi(brandQ)
			typeID, typeErr := strconv.Atoi(typeQ)

			var state *app.UserState

			if (brandOk && brandErr == nil) || (typeOk && typeErr == nil) || (pageOk && pageErr == nil) {
				state = ac.AppState(update.SentFrom().ID)
			} else {
				state = bot.ResetUserState(update, ac)
			}

			switch vars["entity"] {
			case "product":
				conversations.AddProduct(update, ac, state, brandID, typeID, page, msID)
			case "brand":
				conversations.AddBrand(update, ac, state)
			case "type":
				conversations.AddProductType(update, ac, state)
			}
			return
		}

		if vars["entity"] == "product" {
			pID, idOk := queries["pid"]
			id, idErr := strconv.Atoi(pID)

			if !idOk || idErr != nil {
				log.Println(idErr)
				bot.SendText(ac, update, "محصول نامعتبر!")
				return
			}

			switch vars["action"] {
			case "get":
				commands.GetProduct(ac, update, int64(id))
			case "update":
				field, fieldOk := queries["field"]

				if fieldOk && field != "" {
					switch field {
					case "title", "description":
						state := bot.ResetUserState(update, ac)
						state.Data["id"] = pID
						state.Data["field"] = field
						conversations.UpdateProductInfo(update, ac, state)
					case "brand", "type":
						metaId := 0
						brandQ, brandOk := queries["brand"]
						typeQ, typeOk := queries["type"]
						var state *app.UserState

						brandID, brandErr := strconv.Atoi(brandQ)
						typeID, typeErr := strconv.Atoi(typeQ)

						if brandOk && brandErr == nil && brandID != 0 {
							metaId = brandID
							state = ac.AppState(update.SentFrom().ID)
						} else if typeOk && typeErr == nil && typeID != 0 {
							metaId = typeID
							state = ac.AppState(update.SentFrom().ID)
						} else {
							state = bot.ResetUserState(update, ac)

							state.Data["id"] = pID
							state.Data["field"] = field
						}

						conversations.UpdateProductMeta(update, ac, state, metaId, page, msID)
					case "files":
						state := bot.ResetUserState(update, ac)
						state.Data["id"] = pID

						conversations.UpdateProductFiles(update, ac, state)
					}
				} else {
					commands.UpdateProductMenu(ac, update, id)
				}
			case "remove":
				state := bot.ResetUserState(update, ac)
				state.Data["id"] = pID
				conversations.RemoveProduct(update, ac, state)
			case "files":
			}

			return
		}

		if vars["entity"] == "brand" || vars["entity"] == "type" {
			actionText := map[string]string{"remove": "حذف", "update": "ویرایش"}

			categoryText := map[string]string{"brand": "برند", "type": "دسته‌بندی محصول"}

			brandQ, brandOk := queries["brand"]
			typeQ, typeOk := queries["type"]

			brandID, brandErr := strconv.Atoi(brandQ)
			typeID, typeErr := strconv.Atoi(typeQ)

			if vars["entity"] == "brand" && brandOk && brandErr == nil {
				if vars["action"] == "remove" {
					commands.RemoveBrand(ac, update, brandID)
				}
				if vars["action"] == "update" {
					state := bot.ResetUserState(update, ac)
					state.Data["id"] = brandQ
					conversations.UpdateBrand(update, ac, state)
				}
				return
			}

			if vars["entity"] == "type" && typeOk && typeErr == nil {
				if vars["action"] == "remove" {
					commands.RemoveType(ac, update, typeID)
				}
				if vars["action"] == "update" {
					state := bot.ResetUserState(update, ac)
					state.Data["id"] = typeQ
					conversations.UpdateProductType(update, ac, state)

				}
				return
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
		}
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
