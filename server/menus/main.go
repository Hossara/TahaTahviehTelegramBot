package menus

var MainMenu = [][]MenuItem{
	{
		{Path: "/product_list", IsAdmin: false, Name: "لیست محصولات"},
	},
	{
		{Path: "/consultation", IsAdmin: false, Name: "مشاوره تلفنی"},
	},
	{
		{Path: "/remove_product", IsAdmin: true, Name: "حذف محصول"},
		{Path: "/add_product", IsAdmin: true, Name: "افزودن محصول"},
		{Path: "/update_product", IsAdmin: true, Name: "ویرایش محصول"},
	},
	{
		{Path: "/support", IsAdmin: false, Name: "ارتباط‌ با پشتیبانی"},
		{Path: "/faq", IsAdmin: false, Name: "سوالات متداول"},
	},
	{
		{Path: "/edit_faq", IsAdmin: true, Name: "ویرایش سوالات متداول"},
	},
	{
		{Path: "/about", IsAdmin: false, Name: "درباره ما"},
		{Path: "/edit_about", IsAdmin: true, Name: "ویرایش درباره ما"},
	},
}
