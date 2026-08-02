Development
-----------
make setup
make dev
make stop
make clean

Backend
--------
make run
make build
make test
make lint
make fmt
make ai-eval

Database
---------
make migrate-up
make migrate-down
make migrate-create
make seed

Docker
------
make docker-up
make docker-down
make docker-logs

Documentation
--------------
make docs

OpenCode
---------
make ai-check
make ai-docs

Utilities
----------
make help

make prompt

# --- Backend targets ---------------------------------------------------------
BACKEND_DIR := backend

.PHONY: run build test lint fmt ai-eval

run:
	cd $(BACKEND_DIR) && go run ./cmd/api

build:
	cd $(BACKEND_DIR) && go build ./...

test:
	cd $(BACKEND_DIR) && go test -race ./...

lint:
	cd $(BACKEND_DIR) && go vet ./...

fmt:
	cd $(BACKEND_DIR) && gofmt -w .

# Runs the AI evaluation harness against the configured provider
# (M8-003, M4-DEC-012). Requires AI_PROVIDER and AI_API_KEY.
ai-eval:
	cd $(BACKEND_DIR) && go run ./cmd/ai-eval
