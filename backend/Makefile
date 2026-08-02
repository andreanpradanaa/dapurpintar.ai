.PHONY: build test lint vet fmt migrate-up migrate-down itest clean

APP_NAME ?= dapurpintar-api
DB_URL  ?= postgres://dapurpintar:dapurpintar@localhost:5432/dapurpintar?sslmode=disable

build:
	go build -o bin/$(APP_NAME) ./cmd/api

test:
	go test ./...

itest:
	TEST_DATABASE_URL="$(DB_URL)" go test ./... -count=1

lint:
	go vet ./...
	go build ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

migrate-up:
	goose -dir migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DB_URL)" down

run: build
	./bin/$(APP_NAME)

clean:
	rm -rf bin/
