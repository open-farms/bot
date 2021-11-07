.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -trimpath -o bin/bot $(BOT)

.PHONY: lint
lint:
	goimports -w ./