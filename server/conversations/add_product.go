package conversations

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"taha_tahvieh_tg_bot/app"
	productDomain "taha_tahvieh_tg_bot/internal/product/domain"
	psDomain "taha_tahvieh_tg_bot/internal/product_storage/domain"
	storageDomain "taha_tahvieh_tg_bot/internal/storage/domain"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"
	"taha_tahvieh_tg_bot/server/keyboards"
	"taha_tahvieh_tg_bot/server/menus"
)

func AddProduct(update tgbotapi.Update, ac app.App, state *app.UserState, brandID int, typeID, page, prev int) {
	switch state.Step {
	case 0:
		state.Active = true
		state.Conversation = "add_product"

		bot.SendText(ac, update, "نام محصول را بنویسید")
		state.Step = 1

	case 1:
		if update.Message.Text == "" {
			bot.SendText(ac, update, "لطفا یک نام معتبر ارسال کنید!")
			return
		}

		state.Data["name"] = strings.TrimSpace(update.Message.Text)

		bot.SendText(ac, update, "توضیحات محصول را بنویسید")
		state.Step = 2

	case 2:
		if state.Data["description"] == "" {
			if update.Message.Text == "" {
				bot.SendText(ac, update, "لطفا یک متن توضیحات معتبر ارسال کنید!")
				return
			}

			state.Data["description"] = strings.TrimSpace(update.Message.Text)
		}

		if brandID == 0 {
			commands.SelectProductMenu(
				ac, update, "add_product", "brand",
				"لطفا برند محصول را انتخاب کنید",
				page, prev, map[string]string{},
			)
		} else {
			state.Step = 3
			AddProduct(update, ac, state, brandID, typeID, page, prev)
		}
	case 3:
		if brandID == 0 {
			bot.SendText(ac, update, "آیدی برند نامعتبر بود")
			return
		}
		state.Data["brand"] = strconv.Itoa(brandID)

		if typeID == 0 {
			commands.SelectProductMenu(
				ac, update, "add_product", "type",
				"لطفا دسته‌بندی محصول را انتخاب کنید",
				page, prev, map[string]string{
					"brand": strconv.Itoa(brandID),
				},
			)
		} else {
			state.Step = 4
			AddProduct(update, ac, state, brandID, typeID, page, prev)
		}
	case 4:
		if typeID == 0 {
			bot.SendText(ac, update, "آیدی دسته‌بندی نامعتبر بود")
			return
		}

		brandId, err := strconv.Atoi(state.Data["brand"])

		if err != nil {
			bot.SendText(ac, update, "آیدی برند نامعتبر بود")
			return
		}

		id, err := ac.ProductService().CreateProduct(&psDomain.Product{
			UUID:  uuid.New(),
			Title: state.Data["name"],
			Brand: productDomain.Brand{
				ID: productDomain.BrandID(brandId),
			},
			Type: productDomain.ProductType{
				ID: productDomain.ProductTypeID(typeID),
			},
			Description: state.Data["description"],
			Files:       make([]psDomain.File, 0),
		})

		if err != nil {
			log.Println(err)
			bot.SendText(ac, update, "خطا هنگام افزودن محصول")
			return
		}

		state.Data["id"] = strconv.FormatInt(int64(id), 10)

		bot.SendText(ac, update, "فایل های محصول خود را ارسال کنید و در انتها کلمه تمام را ارسال کنید. فرمت های قابل قبول: pdf, png, jpeg")
		state.Step = 5
	case 5:
		pId, err := strconv.Atoi(state.Data["id"])
		if err != nil {
			bot.SendText(ac, update, "محصول نامعتبر است!")
		}

		var fileIDs []storageDomain.FileID

		if len(update.Message.Photo) > 0 || update.Message.Document != nil {
			if len(update.Message.Photo) > 0 {
				bot.SendText(ac, update, "لطفا تا اتمام بارگزاری تمامی فایل ها پیامی ارسال نکنید!")

				photo := update.Message.Photo[len(update.Message.Photo)-1]
				fileURL, _ := ac.Bot().GetFileDirectURL(photo.FileID)
				bot.SendText(ac, update, "درحال آپلود تصویر...")

				id, err := ac.StorageService().UploadFile(&psDomain.File{
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

				fileIDs = append(fileIDs, id)

				bot.SendText(ac, update, "بارگزاری تصویر با موفقیت انجام شد.")
			}

			if update.Message.Document != nil {
				fileID := update.Message.Document.FileID
				fileURL, _ := ac.Bot().GetFileDirectURL(fileID)
				bot.SendText(ac, update, "در حال آپلود فایل ارسال شده...")
				file := update.Message.Document

				id, err := ac.StorageService().UploadFile(&psDomain.File{
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

				fileIDs = append(fileIDs, id)

				bot.SendText(ac, update, "فایل با موفقیت آپلود شد.")
			}

			marshal, err := json.Marshal(fileIDs)

			if err != nil {
				log.Println(err)
				bot.SendText(ac, update, "خطا ذخیره شناسه تصاویر!")
				return
			}

			state.Data["file_ids"] = string(marshal)
			return
		}

		if update.Message.Text != "" {
			msg := tgbotapi.NewMessage(update.FromChat().ID, "محصول با موفقیت افزوده شد!")

			msg.ReplyMarkup = keyboards.InlineKeyboardColumn([]menus.MenuItem{
				{Path: "/menu", IsAdmin: true, Name: "منو اصلی"},
			}, true)

			bot.SendMessage(ac, msg)
			state.Active = false
			ac.DeleteUserState(update.SentFrom().ID)
		}
	}
}
