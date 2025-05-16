package router

import (
	"github.com/NicoNex/echotron/v3"
)

// @route /start
func startHandler(api *echotron.API, m *echotron.Message) {
	api.SendMessage("Welcome! Use /help for options.", m.Chat.ID, nil)
}

// @route /help
func helpHandler(api *echotron.API, m *echotron.Message) {
	api.SendMessage("Try /start or /help!", m.Chat.ID, nil)
}
