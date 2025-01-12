package handlers

import (
	"context"
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
