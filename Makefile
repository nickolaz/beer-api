SHELL:=/bin/bash -O extglob
BINARY=beer-api
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

#go tool command
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/main/main.go cmd/main/routes.go cmd/main/server.go

run:
	go run cmd/main/main.go cmd/main/routes.go cmd/main/server.go

##docker compose
up:
	docker-compose up --build
down:
	docker-compose down --remove-orphans
