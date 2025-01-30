package menus

var MainMenu = [][]MenuItem{
	{
		{Path: "/search", IsAdmin: false, Name: "جستجوی محصول"},
		{Path: "/manage_product", IsAdmin: true, Name: "مدیریت محصولات"},
	},
	{
		{Path: "/manage_brands", IsAdmin: true, Name: "مدیریت برند ها"},
		{Path: "/manage_product_types", IsAdmin: true, Name: "مدیریت دسته‌بندی ها"},
	},
	{
		{Path: "/support", IsAdmin: false, Name: "ارتباط‌ با پشتیبانی"},
	},
	{
		{Path: "/faq", IsAdmin: false, Name: "سوالات متداول"},
	},
	{
		{Path: "/faq_menu", IsAdmin: true, Name: "منو سوالات پرتکرار"},
	},
	{
		{Path: "/edit_about", IsAdmin: true, Name: "ویرایش درباره ما"},
		{Path: "/about", IsAdmin: false, Name: "درباره ما"},
	},
	{
		{Path: "/edit_help", IsAdmin: true, Name: "ویرایش راهنمای ربات"},
		{Path: "/help", IsAdmin: false, Name: "راهنمای ربات"},
	},
}
