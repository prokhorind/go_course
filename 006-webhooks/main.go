package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	ctx := context.Background()
	botToken := os.Getenv("TOKEN")

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set up a webhook on Telegram side
	_ = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://example.com/bot",
		SecretToken: bot.SecretToken(),
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo(ctx)
	fmt.Printf("Webhook Info: %+v\n", info)

	// Create http serve mux
	mux := http.NewServeMux()

	// Get an update channel from webhook.
	// (more on configuration in examples/updates_webhook/main.go)
	updates, _ := bot.UpdatesViaWebhook(ctx, telego.WebhookHTTPServeMux(mux, "/bot", bot.SecretToken()))

	// Start server for receiving requests from the Telegram
	go func() {
		_ = http.ListenAndServe(":443", mux)
	}()

	// Loop through all updates when they came
	for update := range updates {
		if update.Message != nil {
			// Retrieve chat ID
			chatID := update.Message.Chat.ID

			// Call method sendMessage.
			// Send a message to sender with the same text (echo bot).
			// (https://core.telegram.org/bots/api#sendmessage)
			sentMessage, _ := bot.SendMessage(ctx,
				tu.Message(
					tu.ID(chatID),
					update.Message.Text,
				),
			)

			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
