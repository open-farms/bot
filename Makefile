.PHONY: build
build:
	GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" -trimpath -o bin/bot $(BOT)

.PHONY: lint
lint:
	goimports -w ./

.PHONY: push
push:
	push -host pi@10.10.10.10 $(BOT)

.PHONY: broker
broker:
	docker run --rm -v config:/etc/ -p 1883:1883 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqx/emqx:latest

.PHONY: certs
certs:
	openssl genrsa -out config/certs/ca.key 2048
	openssl req -x509 -new -nodes -key config/certs/ca.key -sha256 -days 3650 -out config/certs/ca.pem
	openssl genrsa -out config/certs/emqx.key 2048
	openssl req -new -key ./config/certs/emqx.key -config config/certs/openssl.cnf -out config/certs/emqx.csr
	openssl x509 -req -in ./config/certs/emqx.csr -CA config/certs/ca.pem -CAkey config/certs/ca.key -CAcreateserial -out config/certs/emqx.pem -days 3650 -sha256 -extensions v3_req -extfile config/certs/openssl.cnf
	openssl verify -CAfile config/certs/ca.pem config/certs/emqx.pem