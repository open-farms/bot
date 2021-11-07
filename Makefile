.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -o bin/bot cmd/bot/main.go