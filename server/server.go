package server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"taha_tahvieh_tg_bot/config"
	"taha_tahvieh_tg_bot/server/handlers"

	di "taha_tahvieh_tg_bot/app"
)

func Bootstrap(ctx context.Context, cfg config.Config) error {
	// Init new fiber app
	app := fiber.New()

	// Init cors middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// Init telegram bot client
	bot := InitClient(cfg.Server)
	appContainer := di.MustNewApp(ctx, cfg, bot)

	handlers.Handle(ctx, cfg.Server, appContainer)

	// Start app
	log.Fatal(app.Listen(fmt.Sprintf(":%v", cfg.Server.Port)))
	return nil
}
