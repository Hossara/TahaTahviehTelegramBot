package menus

var SearchProductMenu = [][]MenuItem{
	{
		{Path: "/search/title?page=1", IsAdmin: true, Name: "براساس نام محصول"},
	},
	{
		{Path: "/search/brand?page=1", IsAdmin: true, Name: "براساس برند و دسته‌بندی"},
	},
}

var ManageProductTypeMenu = [][]MenuItem{
	{
		{Path: "/product/type/add", IsAdmin: true, Name: "افزودن دسته‌بندی محصول"},
	},
	{
		{Path: "/product/type/remove?page=1", IsAdmin: true, Name: "حذف دسته‌بندی محصول"},
	},
	{
		{Path: "/product/type/update?page=1", IsAdmin: true, Name: "ویرایش دسته‌بندی محصول"},
	},
}

var ManageBrandMenu = [][]MenuItem{
	{
		{Path: "/product/brand/add", IsAdmin: true, Name: "افزودن برند محصول"},
	},
	{
		{Path: "/product/brand/remove?page=1", IsAdmin: true, Name: "حذف برند محصول"},
	},
	{
		{Path: "/product/brand/update?page=1", IsAdmin: true, Name: "ویرایش برند محصول"},
	},
}
