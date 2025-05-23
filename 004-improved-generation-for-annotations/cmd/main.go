package main

import (
	"github.com/prokhorind/go_course/004-improved-generation-for-annotations/bot"
	"log"
	"os"
)

func main() {

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	engine := bot.NewEchotronBot(token)
	engine.Run()
}
