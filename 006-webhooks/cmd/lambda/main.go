package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"006-webhooks/internal/botlogic"
	"github.com/mymmrac/telego"
)

var bot *telego.Bot

func init() {
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN env var required")
	}

	var err error
	bot, err = botlogic.NewBot(token)
	if err != nil {
		panic(fmt.Sprintf("bot init error: %v", err))
	}
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var update telego.Update
	if err := json.Unmarshal([]byte(req.Body), &update); err != nil {
		fmt.Println("Failed to parse update:", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Use shared update handler
	botlogic.HandleUpdate(ctx, bot, update)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "",
	}, nil
}

// terraform apply -var="telegram_token="
func main() {
	lambda.Start(handler)
}
