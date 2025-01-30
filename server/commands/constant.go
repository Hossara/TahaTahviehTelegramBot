package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var BotCommands = []tgbotapi.BotCommand{
	{Command: "/start", Description: "شروع بات"},
	{Command: "/menu", Description: "منو بات"},
	{Command: "/search", Description: "جستجوی محصولات"},
	{Command: "/support", Description: "ارتباط‌ با پشتیبانی"},
	{Command: "/faq", Description: "سوالات متداول"},
	{Command: "/help", Description: "راهنما ربات"},
	{Command: "/about", Description: "درباره ما"},
}

var AdminCommands = []tgbotapi.BotCommand{
	// Static Replays
	{Command: "/edit_about", Description: "ویرایش پیام درباره من"},
	{Command: "/edit_help", Description: "ویرایش پیام راهنمای ربات"},
	// Manage Product
	{Command: "/manage_product", Description: "مدیریت محصولات"},
	{Command: "/manage_brands", Description: "مدیریت برند ها"},
	{Command: "/manage_product_types", Description: "مدیریت دسته‌بندی ها"},
	// Product
	{Command: "/add_product", Description: "افزودن محصول"},
	{Command: "/remove_product", Description: "حذف محصول"},
	{Command: "/update_product", Description: "ویرایش محصول"},
	// Product Type
	{Command: "/add_product_type", Description: "افزودن دسته‌بندی محصول"},
	{Command: "/remove_product_type", Description: "حذف دسته‌بندی محصول"},
	{Command: "/update_product_type", Description: "ویرایش دسته‌بندی محصول"},
	// Brand
	{Command: "/add_brand", Description: "افزودن برند محصول"},
	{Command: "/remove_brand", Description: "حذف برند محصول"},
	{Command: "/update_brand", Description: "ویرایش برند محصول"},
	// FAQ
	{Command: "/faq_menu", Description: "منو سوال پرتکرار"},
	{Command: "/add_faq", Description: "افزودن سوال پرتکرار"},
	{Command: "/remove_faq", Description: "حذف سوال پرتکرار"},
	{Command: "/update_faq", Description: "ویرایش سوال پرتکرار"},
}
