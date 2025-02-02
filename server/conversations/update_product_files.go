package conversations

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"log"
	"strconv"
	"taha_tahvieh_tg_bot/app"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	psDomain "taha_tahvieh_tg_bot/internal/product_storage/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func UpdateProductFiles(update tgbotapi.Update, ac app.App, state *app.UserState) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "update_product_files"

		bot.SendText(ac, update, "فایل های جدید را ارسال کنید.\nبرای ثبت اطلاعات کلمه تمام را ارسال کنید.\nفرمت های قابل قبول: pdf, png, jpeg")

		pId, err := strconv.Atoi(state.Data["id"])
		if err != nil {
			bot.SendText(ac, update, "محصول نامعتبر است!")
		}

		product, err := ac.ProductService().GetProduct(productDomain.ProductID(pId))

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "خطا هنگام دریافت محصول")
			return
		}

		if product.Files != nil && len(product.Files) > 0 {
			err = ac.StorageService().RemoveAllProductFiles(productDomain.ProductID(pId), product.Files)

			if err != nil {
				log.Println(err)
				bot.SendText(ac, update, "خطا هنگام حذف فایل های قبلی محصول")
				return
			}
		}

		state.Step = 1
	case 1:
		pId, err := strconv.Atoi(state.Data["id"])
		if err != nil {
			bot.SendText(ac, update, "محصول نامعتبر است!")
		}

		if len(update.Message.Photo) > 0 || update.Message.Document != nil {
			if len(update.Message.Photo) > 0 {
				bot.SendText(ac, update, "لطفا تا اتمام بارگزاری تمامی فایل ها پیامی ارسال نکنید!")

				photo := update.Message.Photo[len(update.Message.Photo)-1]
				fileURL, _ := ac.Bot().GetFileDirectURL(photo.FileID)
				bot.SendText(ac, update, "درحال آپلود تصویر...")

				_, err := ac.StorageService().UploadFile(&psDomain.File{
					ProductID:  productDomain.ProductID(pId),
					UUID:       uuid.New(),
					BucketName: "products",
					Path:       "/products/files/",
					Size:       int64(photo.FileSize),
				}, fileURL)

				if err != nil {
					log.Println(err)
					bot.SendText(ac, update, "بارگزاری تصویر با خطا مواجه شد!")
					return
				}

				bot.SendText(ac, update, "بارگزاری تصویر با موفقیت انجام شد.")
			}

			if update.Message.Document != nil {
				fileID := update.Message.Document.FileID
				fileURL, _ := ac.Bot().GetFileDirectURL(fileID)
				bot.SendText(ac, update, "در حال آپلود فایل ارسال شده...")
				file := update.Message.Document

				_, err := ac.StorageService().UploadFile(&psDomain.File{
					ProductID:  productDomain.ProductID(pId),
					UUID:       uuid.New(),
					BucketName: "products",
					Path:       "/products/files/",
					Size:       int64(file.FileSize),
				}, fileURL)

				if err != nil {
					log.Println(err)
					bot.SendText(ac, update, "خطا هنگام آپلود فایل!")
					return
				}

				bot.SendText(ac, update, "فایل با موفقیت آپلود شد.")
			}

			return
		}

		if update.Message.Text != "" {
			msg := tgbotapi.NewMessage(update.FromChat().ID, "محصول با موفقیت ویرایش شد!")

			msg.ReplyMarkup = keyboards.InlineKeyboardColumn([]menus.MenuItem{
				{Path: "/menu", IsAdmin: true, Name: "منو اصلی"},
				{Path: fmt.Sprintf("/product/product/get?pid=%d", pId), IsAdmin: true, Name: "مشاهده محصول"},
			}, true)

			bot.SendMessage(ac, msg)
			state.Active = false
			ac.DeleteUserState(update.SentFrom().ID)
		}
	}
}
