build:
	go build -o ./dist/handler

run:
	go run handler.go

clean:
	rm -rf ./dist