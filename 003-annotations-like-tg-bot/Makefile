# Variables
APP_NAME := myapp
MAIN := main.go

.PHONY: build run deps generate

# Install dependencies
deps:
	go get github.com/NicoNex/echotron/v3

# Run go:generate
generate:
	go generate router/generate_routes.go

# Build the application
build: deps generate
	go build -o $(APP_NAME) $(MAIN)

# Run the application
run: deps generate
	go run $(MAIN)
