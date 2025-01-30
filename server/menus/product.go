package menus

var ManageProductMenu = [][]MenuItem{
	{
		{Path: "/add_product", IsAdmin: true, Name: "افزودن محصول"},
	},
	{
		{Path: "/remove_product", IsAdmin: true, Name: "حذف محصول"},
	},
	{
		{Path: "/update_product", IsAdmin: true, Name: "ویرایش محصول"},
	},
}

var SearchProductMenu = [][]MenuItem{
	{
		{Path: "/search_title", IsAdmin: true, Name: "براساس نام محصول"},
	},
	{
		{Path: "/search_brand", IsAdmin: true, Name: "براساس نام برند"},
	},
	{
		{Path: "/search_type", IsAdmin: true, Name: "براساس نوع محصول"},
	},
}

var ManageProductTypeMenu = [][]MenuItem{
	{
		{Path: "/add_product_type", IsAdmin: true, Name: "افزودن دسته‌بندی محصول"},
	},
	{
		{Path: "/remove_product_type", IsAdmin: true, Name: "حذف دسته‌بندی محصول"},
	},
	{
		{Path: "/update_product_type", IsAdmin: true, Name: "ویرایش دسته‌بندی محصول"},
	},
}

var ManageBrandMenu = [][]MenuItem{
	{
		{Path: "/add_brand", IsAdmin: true, Name: "افزودن برند محصول"},
	},
	{
		{Path: "/remove_brand", IsAdmin: true, Name: "حذف برند محصول"},
	},
	{
		{Path: "/update_brand", IsAdmin: true, Name: "ویرایش برند محصول"},
	},
}
