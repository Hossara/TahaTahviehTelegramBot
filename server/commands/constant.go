package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var BotCommands = []tgbotapi.BotCommand{
	{
		Command:     "/start",
		Description: "شروع بات",
	},
	{
		Command:     "/menu",
		Description: "منو بات",
	},
	{
		Command:     "/product_list",
		Description: "لیست محصولات",
	},
	{
		Command:     "/support",
		Description: "ارتباط‌ با پشتیبانی",
	},
	{
		Command:     "/faq",
		Description: "سوالات متداول",
	},
	{
		Command:     "/help",
		Description: "راهنما ربات",
	},
	{
		Command:     "/about",
		Description: "درباره ما",
	},
}

var AdminCommands = []tgbotapi.BotCommand{
	// Static Replays
	{
		Command:     "/edit_about",
		Description: "ویرایش پیام درباره من",
	}, {
		Command:     "/edit_help",
		Description: "ویرایش پیام راهنمای ربات",
	},
	// Product
	{
		Command:     "/remove_product",
		Description: "حذف محصول",
	},
	{
		Command:     "/add_product",
		Description: "افزودن محصول",
	},
	{
		Command:     "/update_product",
		Description: "ویرایش محصول",
	},
	// FAQ
	{
		Command:     "/faq_menu",
		Description: "منو سوال پرتکرار",
	},
	{
		Command:     "/add_faq",
		Description: "افزودن سوال پرتکرار",
	},
	{
		Command:     "/remove_faq",
		Description: "حذف سوال پرتکرار",
	},
	{
		Command:     "/update_faq",
		Description: "ویرایش سوال پرتکرار",
	},
}
