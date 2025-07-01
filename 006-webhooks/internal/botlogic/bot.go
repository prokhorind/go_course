package botlogic

import (
	"context"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// NewBot creates a new telego.Bot instance with the given token.
func NewBot(token string) (*telego.Bot, error) {
	return telego.NewBot(token, telego.WithDefaultDebugLogger())
}

// HandleUpdate processes an incoming Telegram update.
func HandleUpdate(ctx context.Context, bot *telego.Bot, update telego.Update) {
	if update.Message != nil {
		chatID := update.Message.Chat.ID

		sentMessage, err := bot.SendMessage(ctx,
			tu.Message(
				tu.ID(chatID),
				update.Message.Text,
			),
		)
		if err != nil {
			fmt.Println("SendMessage error:", err)
		} else {
			fmt.Printf("Sent Message: %v\n", sentMessage)
		}
	}
}
