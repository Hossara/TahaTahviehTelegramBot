package menus

var MainMenu = [][]MenuItem{
	{
		{Path: "/search", IsAdmin: false, Name: "جستجوی محصول"},
		{Path: "/manage/product", IsAdmin: true, Name: "افزودن محصول"},
	},
	{
		{Path: "/manage/brands", IsAdmin: true, Name: "مدیریت برند ها"},
		{Path: "/product/product/add", IsAdmin: true, Name: "مدیریت دسته‌بندی ها"},
	},
	{
		{Path: "/support", IsAdmin: false, Name: "ارتباط‌ با پشتیبانی"},
	},
	{
		{Path: "/faq", IsAdmin: false, Name: "سوالات متداول"},
	},
	{
		{Path: "/faq/menu", IsAdmin: true, Name: "منو سوالات پرتکرار"},
	},
	{
		{Path: "/about/update", IsAdmin: true, Name: "ویرایش درباره ما"},
		{Path: "/about", IsAdmin: false, Name: "درباره ما"},
	},
	{
		{Path: "/help/update", IsAdmin: true, Name: "ویرایش راهنمای ربات"},
		{Path: "/help", IsAdmin: false, Name: "راهنمای ربات"},
	},
}
