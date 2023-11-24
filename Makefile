build:
	go build -o ./dist/runner/main

run@nnbot:
	BOT_NAME="tg-bot" go run ./runner/tg-bot.go
run@tgapi:
	BOT_NAME="tg-api" go run ./runner/tg-api

clean:
	rm -rf ./dist