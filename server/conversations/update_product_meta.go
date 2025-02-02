package conversations

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func UpdateProductMeta(update tgbotapi.Update, ac app.App, state *app.UserState, id, page, prev int) {
	var names = map[string]string{
		"brand": "برند",
		"type":  "دسته‌بندی",
	}

	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_product_meta"

		field := state.Data["field"]
		if field == "" {
			bot.SendText(ac, update, "فیلد ویرایش نامشخص")
			return
		}

		if id == 0 {
			commands.SelectProductMenu(
				ac, update, "update_product", field,
				fmt.Sprintf("لطفا %s جدید محصول را انتخاب کنید", names[field]),
				page, prev, map[string]string{
					"pid": state.Data["id"],
				},
			)
		} else {
			state.Data["meta_id"] = strconv.Itoa(id)

			state.Step = 1
			UpdateProductMeta(update, ac, state, id, page, prev)
		}
	case 1:
		pIDQ, pIDOk := state.Data["id"]
		metaIDQ, metaIDOk := state.Data["meta_id"]
		field := state.Data["field"]

		if field == "" {
			bot.SendText(ac, update, "فیلد ویرایش نامشخص")
			return
		}

		metaID, metaIDErr := strconv.Atoi(metaIDQ)
		pID, pIDErr := strconv.Atoi(pIDQ)

		if !pIDOk || pIDErr != nil {
			log.Println(pIDErr)
			bot.SendText(ac, update, "محصول نامعتبر!")
			return
		}

		if !metaIDOk || metaIDErr != nil {
			log.Println(pIDErr)
			bot.SendText(ac, update, fmt.Sprintf("%s نامعتبر!", names[field]))
			return
		}

		err := ac.ProductService().UpdateProduct(productDomain.ProductID(pID), map[string]interface{}{
			fmt.Sprintf("%s_id", field): metaID,
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
