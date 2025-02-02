package conversations

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"taha_tahvieh_tg_bot/app"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

var names = map[string]string{
	"title":       "نام",
	"description": "توضیحات",
}

func UpdateProductInfo(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_product_info"

		field := state.Data["field"]
		if field == "" {
			bot.SendText(ac, update, "فیلد ویرایش نامشخص")
			return
		}

		bot.SendText(ac, update, fmt.Sprintf("%s جدید محصول خود را بنویسید.", names[field]))
		state.Step = 1
	case 1:
		field := state.Data["field"]
		if field == "" {
			bot.SendText(ac, update, "فیلد ویرایش نامشخص")
			return
		}

		if update.Message.Text == "" {
			bot.SendText(ac, update, fmt.Sprintf("لطفا یک %s معتبر ارسال کنید!", names[field]))
			return
		}

		bot.SendText(ac, update, fmt.Sprintf("درحال ویرایش %s محصول...", names[field]))

		pIDQ, pIDOk := state.Data["id"]
		pID, pIDErr := strconv.Atoi(pIDQ)

		if !pIDOk || pIDErr != nil {
			log.Println(pIDErr)
			bot.SendText(ac, update, "محصول نامعتبر!")
			return
		}

		err := ac.ProductService().UpdateProduct(productDomain.ProductID(pID), map[string]interface{}{
			fmt.Sprintf("%s", field): strings.TrimSpace(update.Message.Text),
		})

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "خطا هنگام ویرایش اطلاعات محصول!")
			return
		}

		msg := tgbotapi.NewMessage(update.FromChat().ID, "محصول با موفقیت ویرایش شد!")

		msg.ReplyMarkup = keyboards.InlineKeyboardColumn([]menus.MenuItem{
			{Path: "/menu", IsAdmin: true, Name: "منو اصلی"},
			{Path: fmt.Sprintf("/product/product/get?pid=%d", pID), IsAdmin: true, Name: "مشاهده محصول"},
		}, true)

		bot.SendMessage(ac, msg)

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
