.PHONY: clean

example: clean
	@go run pkg/example/*.go

clean:
	@rm -rf *.png