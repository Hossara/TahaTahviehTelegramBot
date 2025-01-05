package handlers

import (
	"context"
	"log"
	"taha_tahvieh_tg_bot/app"
	"taha_tahvieh_tg_bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Handle(ctx context.Context, cfg config.ServerConfig, ac app.App) {
	/*	userState := cache.CreateUserCache(ctx)
		actionState := cache.CreateActionCache(ctx)
		ctx = context.WithValue(ctx, "user_state", userState)
		ctx = context.WithValue(ctx, "action_state", actionState)*/

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
				}

				if update.Message != nil && update.Message.IsCommand() {
					HandleCommands(update, ac)
				}

				// Handle messages in conversation
				if update.Message != nil && !update.Message.IsCommand() {
					HandleConversations(update, ac)
				}

				// Handle callback queries (from InlineKeyboard)
				if update.CallbackQuery != nil {
					//HandleCallbacks(ctx, update)
				}
			}
		}
	}()
}
