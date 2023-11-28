
OUTPUT_DIR = dist

build-all: clean
	@echo "Building all service..."

	# Build tg-bot
	@cd cmd/tg-bot && go build -o ../../$(OUTPUT_DIR)/tg-bot

	# Build tg-api
	@cd cmd/tg-api && go build -o ../../$(OUTPUT_DIR)/tg-api

run@tg-api:
	@echo "Running 'tg-api'..."
	@go run cmd/tg-api/main.go

run@tg-bot:
	@echo "Running 'tg-bot'..."
	@go run cmd/tg-bot/main.go

clean:
	@echo "Cleaning cache..."
	@rm -rf $(OUTPUT_DIR)
