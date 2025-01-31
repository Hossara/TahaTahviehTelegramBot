run_plain:
	go run cmd/bot/main.go

ROOT_DIR := ./build

NETWORK_NAME=taha-tahvieh-network

DEV_COMPOSE:= \
		-f $(ROOT_DIR)/minio/docker-compose.dev.yaml \
       	-f $(ROOT_DIR)/docker-compose.dev.yaml

PROD_COMPOSE:= \
		-f $(ROOT_DIR)/minio/docker-compose.prod.yaml \
       	-f $(ROOT_DIR)/docker-compose.prod.yaml

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
	docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) up -d

uplog: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) up

# Target to bring down services
down: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) down

build: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) build $(FLAGS) $(CONTAINER)

# Target to view logs
logs: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) logs -f

# Target to rebuild and restart services
rebuild: ensure-network go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(DEV_COMPOSE) up --build

.PHONY: up down logs rebuild

deploy: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) up -d

deploy-down: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) down

deploy-logs: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) logs -f

deploy-log: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) logs -f $(CONTAINER)

deploy-build: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(PROD_COMPOSE) build $(FLAGS) $(CONTAINER)

.PHONY: deploy deploy-down deploy-logs deploy-logs deploy-build