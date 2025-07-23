# export PATH="$PATH:$(go env GOPATH)/bin"

BIN_DIR      := bin
GO           := go
PROTOC       := /opt/homebrew/bin/protoc
COMPOSE_FILE := ./postgres/docker-compose.yml

.PHONY: all proto protoauth protogateway protoalert build-auth build-gateway run-auth run-gateway docker-up docker-down clean integration-test-alert

## ====================== PROTO ==========================

proto: protoauth protogateway-auth protogateway-alert protoalert

protoauth:
	@export PATH="$$PATH:$$($(GO) env GOPATH)/bin"
	$(PROTOC) \
		--proto_path=auth-service/proto \
		--go_out=paths=source_relative:auth-service/proto \
		--go-grpc_out=paths=source_relative:auth-service/proto \
		auth-service/proto/auth.proto
	@echo "✅"

protogateway-auth:
	@export PATH="$$PATH:$$($(GO) env GOPATH)/bin"
	$(PROTOC) \
		--proto_path=gateway/proto-clients/auth \
		--go_out=paths=source_relative:gateway/proto-clients/auth \
		--go-grpc_out=paths=source_relative:gateway/proto-clients/auth \
		gateway/proto-clients/auth/auth.proto
	@echo "✅"

protogateway-alert:
	@export PATH="$$PATH:$$($(GO) env GOPATH)/bin"
	$(PROTOC) \
		--proto_path=gateway/proto-clients/alert \
		--go_out=paths=source_relative:gateway/proto-clients/alert \
		--go-grpc_out=paths=source_relative:gateway/proto-clients/alert \
		gateway/proto-clients/alert/alert.proto
	@echo "✅"

protoalert:
	@export PATH="$$PATH:$$($(GO) env GOPATH)/bin"
	$(PROTOC) \
		--proto_path=alert-service/proto \
		--go_out=paths=source_relative:alert-service/proto \
		--go-grpc_out=paths=source_relative:alert-service/proto \
		alert-service/proto/alert.proto
	@echo "✅"

## ====================== BUILD ==========================

build-auth:
	$(GO) build -o $(BIN_DIR)/auth ./auth-service/cmd/main.go
	@echo "✅"

build-gateway:
	$(GO) build -o $(BIN_DIR)/gateway ./gateway/cmd/main.go
	@echo "✅"

build-alert:
	$(GO) build -o $(BIN_DIR)/alert ./alert-service/cmd/main.go
	@echo "✅"

build-watcher:
	$(GO) build -o $(BIN_DIR)/watcher ./price-watcher/cmd/main.go
	@echo "✅"

build-notification:
	$(GO) build -o $(BIN_DIR)/notification ./notification-service/cmd/main.go
	@echo "✅"

## ====================== RUN ==========================

run-auth: build-auth
	./$(BIN_DIR)/auth

run-gateway: build-gateway
	./$(BIN_DIR)/gateway

run-alert: build-alert
	./$(BIN_DIR)/alert

run-watcher: build-watcher
	./$(BIN_DIR)/watcher

run-notification: build-notification
	./$(BIN_DIR)/notification

## ====================== TESTS ==========================

integration-test-alert:
	go test ./alert-service/internal/integration/... -v

integration-test-auth:
	go test ./auth-service/internal/integration/... -v

integration-test-watcher:
	go test ./price-watcher/internal/integration/... -v

integration-test-all: integration-test-alert integration-test-auth integration-test-watcher

test-all:
	go test ./... -v
	@echo "✅ All tests complete"

## ====================== DOCKER ==========================

docker-up:
	docker-compose -f $(COMPOSE_FILE) up -d
	@echo "🚀 Docker Compose UP"

docker-down:
	docker-compose -f $(COMPOSE_FILE) down -v
	@echo "Compose DOWN"

## ====================== CLEAN ==========================

clean:
	rm -rf $(BIN_DIR)
	@echo "Cleaned binaries"
