package menus

var MainMenu = [][]MenuItem{
	{
		{Path: "/product_list", IsAdmin: false, Name: "لیست محصولات"},
	},
	{
		{Path: "/remove_product", IsAdmin: true, Name: "حذف محصول"},
		{Path: "/update_product", IsAdmin: true, Name: "ویرایش محصول"},
		{Path: "/add_product", IsAdmin: true, Name: "افزودن محصول"},
	},
	{
		{Path: "/support", IsAdmin: false, Name: "ارتباط‌ با پشتیبانی"},
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
