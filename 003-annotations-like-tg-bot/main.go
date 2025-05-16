package main

//go:generate go run router/generate_routes.go
import (
	"github.com/prokhorind/go_course/003-annotations-like-tg-bot/router"
	"strings"

	"github.com/NicoNex/echotron/v3"
)

const botToken = "YOU_TOKEN"

func main() {
	api := echotron.NewAPI(botToken)

	for update := range echotron.PollingUpdates(botToken) {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}
		text := strings.Split(update.Message.Text, " ")[0]
		if handler, ok := router.GetHandler(text); ok {
			handler(&api, update.Message)
		} else {
			api.SendMessage("Unknown command. Try /help", update.ChatID(), nil)
		}
	}
}
