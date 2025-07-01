# Telegram Ping Bot (Webhook) — Server & AWS Lambda

This project implements a Telegram bot that echoes messages using the [telego](https://github.com/mymmrac/telego) library. It supports two deployment modes:

- **Standalone HTTP Server**: Runs as a webhook server that listens for Telegram updates.
- **AWS Lambda Function**: Runs as a Lambda behind API Gateway with a static HTTPS webhook URL.

---

## Project Structure

```
yourproject/
├── bin/                       # Build outputs (binaries, zipped lambda)
├── cmd/
│   ├── lambda/                # Lambda entrypoint
│   └── server/                # Standalone server entrypoint
├── internal/
│   └── botlogic/              # Shared bot logic & update handler
├── terraform/                 # Terraform config for AWS deployment
├── Makefile                   # Build commands
├── go.mod
└── go.sum
```

---

## Prerequisites

- Go 1.20+ installed
- AWS CLI configured with appropriate permissions
- Terraform installed (version 1.0+ recommended)
- Telegram Bot Token (get it from [BotFather](https://t.me/BotFather))

---

## Local Development (Standalone Server)

1. Set your Telegram bot token:

```bash
export TOKEN=your_telegram_bot_token
```

2. Run ngrok
```
ngrok http 8080
ngrok http --url=obliging-precisely-dove.ngrok-free.app 8080
```

3. Set your  webhook URL to your server’s HTTPS in server/main.go:

```
cmd/server/main.go 
err = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "%YOUR_NGROK_URL%/bot",
		SecretToken: bot.SecretToken(),
	})
```

4. Build and run the server:

```bash
make build-server
./bin/server
```




---

## AWS Lambda Deployment with Terraform

### 1. Build and package the Lambda binary

```bash
make build-and-zip-lambda
```

This will produce `bin/lambda.zip` containing the executable named `bootstrap` required for AWS Lambda custom runtime (`provided.al2`).

### 2. Configure Terraform variables

Edit `terraform/terraform.tfvars` (create it if needed, and do **not** commit it if it contains secrets):

```hcl
telegram_token = "your_telegram_bot_token_here"
```

### 3. Deploy with Terraform

```bash
cd terraform
terraform init
terraform apply
```

### 4. Get the API Gateway URL

After deployment, Terraform will output the API Gateway invoke URL. The webhook URL for Telegram is:

```
<API_GATEWAY_INVOKE_URL>/bot
```

Set this URL as the webhook for your Telegram bot:

```bash
https://stackoverflow.com/questions/42554548/how-to-set-telegram-bot-webhook
```

---

## Cleaning up

To remove build artifacts:

```bash
make clean
```

---

## Notes

- The Lambda uses **custom runtime `provided.al2`**, so the binary must be named `bootstrap`.
- The Telegram bot logic is shared between both deployments in `internal/botlogic`.
- Function Url is configured to forward `POST /bot` requests to the Lambda.

---

## Troubleshooting

- If your Lambda fails to execute, check CloudWatch Logs.
- Make sure your Telegram webhook URL is accessible and uses HTTPS.
- Verify the environment variable `TOKEN` is correctly set in Lambda via Terraform.

---

