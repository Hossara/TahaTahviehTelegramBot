package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"taha_tahvieh_tg_bot/config"
	"taha_tahvieh_tg_bot/server"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	fmt.Println("Starting Bot...")

	err := server.Bootstrap(context.Background(), c)

	if err != nil {
		return
	}

	fmt.Println("Bot stopped.")
}
