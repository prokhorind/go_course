APP_NAME := telegram-bot
MAIN_FILE := ./cmd/main.go
ROUTE_GEN := ./scripts/generate_routes.go -handler-dir=./handler -output-file=handler/generated_routes.go -output-pkg=handler
.PHONY: all tidy fmt build run generate-routes

all: tidy fmt generate-routes build

tidy:
	go mod tidy

fmt:
	go fmt ./...

generate-routes:
	go run $(ROUTE_GEN)

build: generate-routes
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

run: generate-routes
	go run $(MAIN_FILE)
