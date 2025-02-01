package menus

var FaqMenu = [][]MenuItem{
	{
		{Path: "/faq/add", IsAdmin: true, Name: "افزودن سوال پرتکرار"},
	},
	{
		{Path: "/faq/menu/update", IsAdmin: true, Name: "ویرایش سوال پرتکرار"},
		{Path: "/faq/menu/remove", IsAdmin: true, Name: "حذف سوال پرتکرار"},
	},
}
