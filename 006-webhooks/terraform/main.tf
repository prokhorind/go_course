provider "aws" {
  region = "eu-central-1"
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
  runtime          = "provided.al2"  # <- change here
  role             = aws_iam_role.lambda_role.arn
  source_code_hash = filebase64sha256("../bin/lambda.zip")

  environment {
    variables = {
      TOKEN = var.telegram_token
    }
  }
}


# Create HTTP API Gateway (HTTP API, v2)
resource "aws_apigatewayv2_api" "http_api" {
  name          = "ping-http-api"
  protocol_type = "HTTP"
}

# Lambda integration for API Gateway
resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id                 = aws_apigatewayv2_api.http_api.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.ping.arn
  integration_method     = "POST"
  payload_format_version = "2.0"
}

# Default route to Lambda integration
resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.http_api.id
  route_key = "POST /bot"    # Telegram webhook endpoint path
  target    = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

# Deployment
resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.http_api.id
  name        = "$default"  # default stage for HTTP APIs
  auto_deploy = true
}

# Permission for API Gateway to invoke Lambda
resource "aws_lambda_permission" "apigw_invoke" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.ping.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.http_api.execution_arn}/*/*"
}

output "lambda_function_name" {
  value = aws_lambda_function.ping.function_name
}

output "api_gateway_invoke_url" {
  value = aws_apigatewayv2_stage.default.invoke_url
}
