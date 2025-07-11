provider "aws" {
  region = "eu-central-1" # Frankfurt
}

# IAM Role for Lambda
resource "aws_iam_role" "lambda_role" {
  name = "ping_lambda_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

# Attach basic execution policy
resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# Lambda function
resource "aws_lambda_function" "ping" {
  filename         = "../bin/lambda.zip"
  function_name    = "ping"
  handler          = "bootstrap"     # for provided.al2 custom runtime, handler is bootstrap
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_role.arn
  source_code_hash = filebase64sha256("../bin/lambda.zip")

  environment {
    variables = {
      TOKEN = var.telegram_token
    }
  }
}

resource "aws_lambda_function_url" "ping_url" {
  function_name      = aws_lambda_function.ping.function_name
  authorization_type = "NONE" # Public access for Telegram webhook
}

output "webhook_url" {
  value = "${aws_lambda_function_url.ping_url.function_url}bot"
}
