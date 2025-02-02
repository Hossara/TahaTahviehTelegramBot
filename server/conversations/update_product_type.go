package conversations

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func UpdateProductType(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_product_type"

		bot.SendText(ac, update, "نام جدید دسته‌بندی را بنویسید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک نام معتبر ارسال کنید!")
			return
		}

		state.Data["name"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "توضیحات جدید دسته‌بندی را بنویسید")
		state.Step = 2
	case 2:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن توضیحات معتبر ارسال کنید!")
			return
		}

		bot.SendText(ac, update, fmt.Sprintf("در حال ویرایش دسته‌بندی %s...", state.Data["name"]))

		typeQ, typeOk := state.Data["id"]
		typeID, typeErr := strconv.Atoi(typeQ)

		if !typeOk || typeErr != nil {
			log.Println(typeErr)
			bot.SendText(ac, update, "دسته‌بندی نامعتبر!")
			return
		}

		err := ac.ProductService().UpdateProductType(domain.ProductTypeID(typeID), map[string]interface{}{
			"title":       state.Data["name"],
			"description": strings.TrimSpace(update.Message.Text),
		})

		if err != nil {
			log.Println("Error while update product type", err)
			bot.SendText(ac, update, "خطا هنگام ویرایش دسته‌بندی!")
			return
		}

		msg := tgbotapi.NewMessage(
			update.FromChat().ID,
			fmt.Sprintf("دسته‌بندی %s با موفقیت ویرایش شد.", state.Data["name"]),
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
