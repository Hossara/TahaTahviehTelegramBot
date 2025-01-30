package handlers

import (
	"context"
	"fmt"
	"log"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/config"
	"taha_tahvieh_tg_bot/pkg/bot"
	"taha_tahvieh_tg_bot/server/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMenu(update tgbotapi.Update, ac app.App) {
	botCommands := commands.BotCommands

	// Commands for admins & super admins
	if bot.IsSuperRole(update, ac) {
		botCommands = append(botCommands, commands.AdminCommands...)
	}

	cmdCfg := tgbotapi.NewSetMyCommands(botCommands...)

	bot.SendRequest(ac, cmdCfg)
}

func checkUserSubscription(bot *tgbotapi.BotAPI, channel string, userID int64) (bool, error) {
	chatMemberConfig := tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			SuperGroupUsername: channel,
			UserID:             userID,
		},
	}

	member, err := bot.GetChatMember(chatMemberConfig)

	if err != nil {
		return false, err
	}

	return member.Status == "member" || member.Status == "administrator" || member.Status == "creator", nil
}

func Handle(ctx context.Context, cfg config.ServerConfig, ac app.App) {
	u := tgbotapi.NewUpdate(int(cfg.NewUpdateOffset))
	u.Timeout = int(cfg.BotTimeout)

	updates := ac.Bot().GetUpdatesChan(u)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Received shutdown signal, stopping update processing...")
				return

			case update, ok := <-updates:
				if !ok {
					log.Println("Updates channel closed.")
					return
				}

				channel := ac.Config().Constants.Channel

				sub, err := checkUserSubscription(ac.Bot(), channel, update.SentFrom().ID)

				if !sub || err != nil {
					bot.SendText(
						ac, update,
						fmt.Sprintf("برای استفاده از بات، ابتدا در کانال %s عضو شوید.", channel),
					)
					continue
				}

				// Handle command menu
				handleMenu(update, ac)

				// Handle commands
				if update.Message != nil && update.Message.IsCommand() {
					HandleCommands(update, ac)
				}

				// Handle messages in conversation
				if update.Message != nil && !update.Message.IsCommand() {
					HandleConversations(update, ac)
				}

				// Handle callback queries (from InlineKeyboard)
				if update.CallbackQuery != nil {
					HandleCallbacks(update, ac)
				}
			}
		}
	}()
}
