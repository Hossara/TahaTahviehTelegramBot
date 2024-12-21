package server

import (
	"context"
	"golang.org/x/net/proxy"
	"log"
	"net"
	"net/http"
	"taha_tahvieh_tg_bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitClient(cfg config.ServerConfig) *tgbotapi.BotAPI {
	// Set up SOCKS5 proxy
	if cfg.Proxy != "" {
		// Create socks 5 connection
		dialer, err := proxy.SOCKS5("tcp", cfg.Proxy, &proxy.Auth{
			User:     "",
			Password: "",
		}, proxy.Direct)

		if err != nil {
			log.Fatalf("failed to connect to proxy: %v", err)
			return nil
		}

		// Use it in http transport layer
		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}

		// Create HTTP client with the proxy
		client := &http.Client{Transport: transport}
		bot, err := tgbotapi.NewBotAPIWithClient(cfg.Token, tgbotapi.APIEndpoint, client)

		if err != nil {
			log.Panic(err)
			return nil
		}

		return bot
	} else {
		// Create telegram bot instance without proxy
		bot, err := tgbotapi.NewBotAPI(cfg.Token)

		if err != nil {
			log.Panic(err)
			return nil
		}

		return bot
	}
}
