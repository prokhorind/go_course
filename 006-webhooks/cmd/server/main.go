package main

import (
	"006-webhooks/internal/botlogic"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
)

// ngrok http --url=obliging-precisely-dove.ngrok-free.app 8080
func main() {
	ctx := context.Background()
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN env var required")
	}

	bot, err := botlogic.NewBot(token)
	if err != nil {
		log.Fatalf("bot init error: %v", err)
	}

	// Set webhook
	err = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://obliging-precisely-dove.ngrok-free.app/bot",
		SecretToken: bot.SecretToken(),
	})
	if err != nil {
		log.Fatalf("set webhook error: %v", err)
	}

	mux := http.NewServeMux()

	updates, err := bot.UpdatesViaWebhook(ctx, telego.WebhookHTTPServeMux(mux, "/bot", bot.SecretToken()))
	if err != nil {
		log.Fatalf("updates webhook error: %v", err)
	}

	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("http listen error: %v", err)
		}
	}()

	for update := range updates {
		botlogic.HandleUpdate(ctx, bot, update)
	}
}
