package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
	"taha_tahvieh_tg_bot/server/conversations"
	"taha_tahvieh_tg_bot/server/menus"
)

func HandleCommands(update tgbotapi.Update, ac app.App) {
	switch update.Message.Command() {
	// -------------------- General
	case "start":
		commands.Start(ac, update)
	case "about":
		commands.About(ac, update)
	case "menu":
		commands.Menu(ac, update)
	case "support":
		commands.Support(ac, update)
	case "help":
		commands.Help(ac, update)
	case "search":
		commands.SearchProductMenu(ac, update)

	// -------------------- Manage Product
	case "manage_product":
		commands.ProductManagementMenu(ac, update, menus.ManageProductMenu)
	case "manage_brands":
		commands.ProductManagementMenu(ac, update, menus.ManageBrandMenu)
	case "manage_product_types":
		commands.ProductManagementMenu(ac, update, menus.ManageProductTypeMenu)

	// -------------------- Products
	case "add_product":
	case "remove_product":
	case "update_product":

	// -------------------- Product Types
	case "add_product_type":
	case "remove_product_type":
	case "update_product_type":

	// -------------------- Brand
	case "add_brand":
	case "remove_brand":
	case "update_brand":

	// -------------------- FAQ
	case "faq":
		commands.FaqList(ac, update)

	case "faq_menu":
		commands.FaqMenu(ac, update)

	case "add_faq":
		state := bot.ResetUserState(update, ac)
		conversations.AddFaq(update, ac, state)

	case "remove_faq":
		commands.RemoveFaqMenu(ac, update)

	case "update_faq":
		commands.UpdateFaqMenu(ac, update)

	// -------------------- General Conversations
	case "edit_about":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateAbout(update, ac, state)

	case "edit_help":
		state := bot.ResetUserState(update, ac)
		conversations.UpdateHelp(update, ac, state)
	}
	return
}
