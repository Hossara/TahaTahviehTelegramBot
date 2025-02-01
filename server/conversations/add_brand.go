package conversations

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/internal/product/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
)

func AddBrand(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "add_brand"

		bot.SendText(ac, update, "نام برند جدید را بنویسید")
		state.Step = 1
	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک نام معتبر ارسال کنید!")
			return
		}

		state.Data["name"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "توضیحات برند را بنویسید")
		state.Step = 2
	case 2:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک متن توضیحات معتبر ارسال کنید!")
			return
		}

		bot.SendText(ac, update, fmt.Sprintf("در حال افزودن برند %s...", state.Data["name"]))

		err := ac.ProductService().CreateBrand(&domain.Brand{
			Title:       state.Data["name"],
			Description: strings.TrimSpace(update.Message.Text),
		})

		if err != nil {
			log.Println("Error while insert new brand", err)
			bot.SendText(ac, update, "خطا هنگام افزودن برند جدید!")
			return
		}

		bot.SendText(ac, update, fmt.Sprintf("برند %s با موفقیت اضافه شد.", state.Data["name"]))

		state.Active = false
		ac.DeleteUserState(update.SentFrom().ID)
	}
}
