# Project metadata
APP_NAME = go-layered-app
BIN_DIR = bin

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "🔨 Building..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd

# Run the application
.PHONY: run
run: build
	@echo "🚀 Running..."
	./$(BIN_DIR)/$(APP_NAME)

# Run all tests
.PHONY: test
test:
	@echo "🧪 Running tests..."
	go test ./... -v

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "📊 Running tests with coverage..."
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "🔍 See coverage.html"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning..."
	rm -rf $(BIN_DIR) coverage.out coverage.html

# Format code
.PHONY: fmt
fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...
