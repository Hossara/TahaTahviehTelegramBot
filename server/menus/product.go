package menus

var SearchProductMenu = [][]MenuItem{
	{
		{Path: "/search/title?page=1", IsAdmin: true, Name: "براساس نام محصول"},
	},
	{
		{Path: "/search/brand?page=1", IsAdmin: true, Name: "براساس برند و دسته‌بندی"},
	},
}

var ManageProductMenu = [][]MenuItem{
	{
		{Path: "/product/menu/product/add", IsAdmin: true, Name: "افزودن محصول"},
	},
	{
		{Path: "/product/menu/product/remove?page=1", IsAdmin: true, Name: "حذف محصول"},
	},
	{
		{Path: "/product/menu/product/update?page=1", IsAdmin: true, Name: "ویرایش محصول"},
	},
}

var ManageProductTypeMenu = [][]MenuItem{
	{
		{Path: "/product/menu/type/add", IsAdmin: true, Name: "افزودن دسته‌بندی محصول"},
	},
	{
		{Path: "/product/menu/type/remove?page=1", IsAdmin: true, Name: "حذف دسته‌بندی محصول"},
	},
	{
		{Path: "/product/menu/type/update?page=1", IsAdmin: true, Name: "ویرایش دسته‌بندی محصول"},
	},
}

var ManageBrandMenu = [][]MenuItem{
	{
		{Path: "/product/menu/brand/add", IsAdmin: true, Name: "افزودن برند محصول"},
	},
	{
		{Path: "/product/menu/brand/remove?page=1", IsAdmin: true, Name: "حذف برند محصول"},
	},
	{
		{Path: "/product/menu/brand/update?page=1", IsAdmin: true, Name: "ویرایش برند محصول"},
	},
}
