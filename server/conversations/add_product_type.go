package conversations

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func AddProductType(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "add_product_type"

		bot.SendText(ac, update, "نام دسته‌بندی جدید را بنویسید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک نام معتبر ارسال کنید!")
			return
		}

		state.Data["name"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "توضیحات دسته‌بندی را بنویسید")
		state.Step = 2
	case 2:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن توضیحات معتبر ارسال کنید!")
			return
		}

		bot.SendText(ac, update, fmt.Sprintf("در حال افزودن دسته‌بندی %s...", state.Data["name"]))

		err := ac.ProductService().CreateProductType(&domain.ProductType{
			Title:       state.Data["name"],
			Description: strings.TrimSpace(update.Message.Text),
		})

		if err != nil {
			log.Println("Error while insert new product type", err)
			bot.SendText(ac, update, "خطا هنگام افزودن دسته‌بندی جدید!")
			return
		}

		msg := tgbotapi.NewMessage(
			update.FromChat().ID,
			fmt.Sprintf("دسته‌بندی %s با موفقیت اضافه شد.", state.Data["name"]),
		)

		msg.ReplyMarkup = keyboards.InlineKeyboardColumn([]menus.MenuItem{
			{Path: "/manage/product_types", IsAdmin: true, Name: "مدیریت دسته‌بندی ها"},
			{Path: "/menu", IsAdmin: true, Name: "منو اصلی"},
		}, true)

		bot.SendMessage(ac, msg)

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
