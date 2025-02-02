package conversations

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func RemoveProduct(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "remove_product"

		bot.SendText(ac, update, "آیا از حذف محصول اطمینان دارید؟ (بلی/خیر)")
		state.Step = 1
	case 1:
		if update.Message.Text == "" || update.Message.Text == "خیر" {
			bot.SendText(ac, update, "عملیات لغو شد.")

			state.Active = false
			ac.DeleteUserState(update.SentFrom().ID)
			return
		}

		id, err := strconv.Atoi(state.Data["id"])

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "محصول نامعتبر!")
			return
		}

		bot.SendText(ac, update, "درحال حذف...")

		product, err := ac.ProductService().GetProduct(productDomain.ProductID(id))

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "خطا هنگام دریافت محصول")
			return
		}

		err = ac.ProductService().DeleteProduct(productDomain.ProductID(id), product.Files)

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "خطا هنگام خذف محصول!")
			return
		}

		msg := tgbotapi.NewMessage(update.FromChat().ID, "محصول با موفقیت حذف شد!")

		msg.ReplyMarkup = keyboards.InlineKeyboardColumn([]menus.MenuItem{
			{Path: "/menu", IsAdmin: true, Name: "منو اصلی"},
		}, true)

		bot.SendMessage(ac, msg)
	}
}
