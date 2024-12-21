run_plain:
	go run cmd/bot/main.go


ROOT_DIR := ./build

NETWORK_NAME=taha-tahvieh-network

ensure-network:
	@echo "Ensuring project network is exists..."
	@if [ -z "$$(docker network ls --filter name=$(NETWORK_NAME) --format '{{.Name}}')" ]; then \
		echo "Network $(NETWORK_NAME) does not exist. Creating..."; \
		docker network create $(NETWORK_NAME); \
	fi

go-mod-vendor:
	@echo "Running 'go mod vendor' to sync dependencies..."
	@go mod tidy
	@go mod vendor

# Default target to bring up services
up: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) up -d

# Target to bring down services
down: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) down

CONTAINER ?= all
FLAGS ?=
build: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) build $(FLAGS) $(CONTAINER)

# Target to view logs
logs: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) logs -f

# Target to rebuild and restart services
rebuild: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) up --build

.PHONY: up down logs rebuild
