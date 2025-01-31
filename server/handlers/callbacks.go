package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"slices"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
	"taha_tahvieh_tg_bot/server/conversations"
	"taha_tahvieh_tg_bot/server/menus"
)

func HandleCallbacks(update tgbotapi.Update, ac app.App) {
	action := update.CallbackQuery.Data
	//chatID := update.CallbackQuery.Message.Chat.ID

	switch {
	// -------------------- General
	case action == "/about":
		commands.About(ac, update)

	case action == "/menu":
		commands.Menu(ac, update)

	case action == "/support":
		commands.Support(ac, update)

	case action == "/help":
		commands.Help(ac, update)

	// -------------------- Search
	case action == "/search":
		commands.SearchProductMenu(ac, update)

	case action == "/search_title":
	case action == "/search_brand":
		commands.SelectProductMenu(ac, update, "brand", "محصولات هر برند")
	case action == "/search_type":
		commands.SelectProductMenu(ac, update, "type", "محصولات هر دسته‌بندی")

	// -------------------- Product Management
	case action == "/manage_product":
		commands.ProductManagementMenu(ac, update, menus.ManageProductMenu)
	case action == "/manage_brands":
		commands.ProductManagementMenu(ac, update, menus.ManageBrandMenu)
	case action == "/manage_product_types":
		commands.ProductManagementMenu(ac, update, menus.ManageProductTypeMenu)

	// -------------------- Products
	case action == "/add_product":
	case action == "/remove_product":
	case action == "/update_product":

	// -------------------- FAQ
	case action == "/faq":
		commands.FaqList(ac, update)

	case action == "/faq_menu":
		commands.FaqMenu(ac, update)

	case action == "/add_faq":
		state := bot.ResetUserState(update, ac)
		conversations.AddFaq(update, ac, state)

	case action == "/remove_faq":
		commands.RemoveFaqMenu(ac, update)

	case action == "/update_faq":
		commands.UpdateFaq(ac, update)

	// -------------------- General Conversations
	case action == "/edit_about":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateAbout(update, ac, state)

	case action == "/edit_help":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateHelp(update, ac, state)
	}

	// Check Path Variables
	actionPath := strings.Split(action, "/")

	switch {
	// Get Exact FAQ
	case slices.Contains(actionPath, "get_faq") && actionPath[2] != "":
		commands.Faq(ac, update, actionPath[2])

	// Delete Exact FAQ Confirmations Question
	case slices.Contains(actionPath, "q_r_faq") && actionPath[2] != "":
		commands.QuestionRemoveFaq(ac, update, actionPath[2])

	// Delete Exact FAQ
	case slices.Contains(actionPath, "del_faq") && actionPath[2] != "":
		commands.RemoveFaq(ac, update, actionPath[2])

	// Update Exact FAQ
	case slices.Contains(actionPath, "update_faq") && len(actionPath) > 2 && actionPath[2] != "":
		state := bot.ResetUserState(update, ac)
		state.Data["id"] = actionPath[2]
		conversations.UpdateFaq(update, ac, state)
	}
}
