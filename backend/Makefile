SHELL := /bin/bash

.PHONY: help run test build tidy seed migrate-install migrate-up migrate-down migrate-create dev

help:
	@echo "Dapur Pintar backend"
	@echo ""
	@echo "Targets:"
	@echo "  run         - start the server (auto-migrate, auto-seed)"
	@echo "  test        - run unit + handler tests"
	@echo "  build       - build the binary"
	@echo "  tidy        - go mod tidy"
	@echo "  seed        - load data/recipes.json into the DB"
	@echo "  migrate-up  - apply all pending migrations"
	@echo "  migrate-down - roll back one migration"
	@echo "  dev         - start postgres in docker, wait, migrate, run"

run:
	go run ./cmd/server

test:
	go test ./... -v

build:
	go build -o bin/server ./cmd/server

tidy:
	go mod tidy

seed:
	go run ./cmd/seed

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up

migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down 1

migrate-create:
	@read -p "migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

dev:
	docker compose up -d postgres
	@echo "Waiting for postgres to be ready…"
	@until docker compose exec -T postgres pg_isready -U dapur >/dev/null 2>&1; do sleep 1; done
	@echo "Running migrations…"
	@DATABASE_URL=postgres://dapur:dapur_dev@localhost:5432/dapur?sslmode=disable go run ./cmd/server
