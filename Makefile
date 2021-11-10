.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -trimpath -o bin/bot $(BOT)

.PHONY: lint
lint:
	goimports -w ./

.PHONY: push
push:
	push -host pi@10.10.10.10 ./_examples/pubsub/subscribe/...

.PHONY: broker
broker:
	 docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqx/emqx 