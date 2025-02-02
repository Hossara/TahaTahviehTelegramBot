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

func UpdateBrand(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_brand"

		bot.SendText(ac, update, "نام جدید برند را بنویسید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک نام معتبر ارسال کنید!")
			return
		}

		state.Data["name"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "توضیحات جدید برند را بنویسید")
		state.Step = 2
	case 2:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن توضیحات معتبر ارسال کنید!")
			return
		}

		bot.SendText(ac, update, fmt.Sprintf("در حال ویرایش برند %s...", state.Data["name"]))

		brandQ, brandOk := state.Data["id"]
		brandID, brandErr := strconv.Atoi(brandQ)

		if !brandOk || brandErr != nil {
			log.Println(brandErr)
			bot.SendText(ac, update, "برند نامعتبر!")
			return
		}

		err := ac.ProductService().UpdateBrand(domain.BrandID(brandID), map[string]interface{}{
			"title":       state.Data["name"],
			"description": strings.TrimSpace(update.Message.Text),
		})

		if err != nil {
			log.Println("Error while update brand", err)
			bot.SendText(ac, update, "خطا هنگام ویرایش برند!")
			return
		}

		msg := tgbotapi.NewMessage(
			update.FromChat().ID,
			fmt.Sprintf("برند %s با موفقیت ویرایش شد.", state.Data["name"]),
		)

		msg.ReplyMarkup = keyboards.InlineKeyboardColumn([]menus.MenuItem{
			{Path: "/manage/brands", IsAdmin: true, Name: "مدیریت برند ها"},
			{Path: "/menu", IsAdmin: true, Name: "منو اصلی"},
		}, true)

		bot.SendMessage(ac, msg)

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
