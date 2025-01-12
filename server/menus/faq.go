package menus

var FaqMenu = [][]MenuItem{
	{
		{Path: "/add_faq", IsAdmin: true, Name: "افزودن سوال پرتکرار"},
	},
	{
		{Path: "/remove_faq", IsAdmin: true, Name: "حذف سوال پرتکرار"},
	},
	{
		{Path: "/update_faq", IsAdmin: true, Name: "ویرایش سوال پرتکرار"},
	},
}
