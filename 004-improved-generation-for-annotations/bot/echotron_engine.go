package bot

import (
	"github.com/prokhorind/go_course/004-improved-generation-for-annotations/handler"
	"log"
	"strings"

	"github.com/NicoNex/echotron/v3"
)

type EchotronContext struct {
	API     *echotron.API
	Message *echotron.Message
}

func (c *EchotronContext) SendMessage(text string) {
	c.API.SendMessage(text, c.Message.Chat.ID, nil)
}

func (c *EchotronContext) ChatID() int64 {
	return c.Message.Chat.ID
}

type EchotronBot struct {
	token string
	api   *echotron.API
}

func NewEchotronBot(token string) handler.BotEngine {
	api := echotron.NewAPI(token)
	return &EchotronBot{
		token: token,
		api:   &api,
	}
}

func (b *EchotronBot) Run() {
	log.Println("🤖 Echotron bot is running...")

	for update := range echotron.PollingUpdates(b.token) {
		if update.Message == nil {
			continue
		}

		text := strings.TrimSpace(update.Message.Text)
		if text == "" || !strings.HasPrefix(text, "/") {
			continue
		}

		cmd := strings.Fields(text)[0]
		ctx := &EchotronContext{
			API:     b.api,
			Message: update.Message,
		}

		if handlerFunc, ok := handler.GetHandler(cmd); ok {
			handlerFunc(ctx)
		} else {
			ctx.SendMessage("Unknown command 🤷")
		}
	}
}
