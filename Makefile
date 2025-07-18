# export PATH="$PATH:$(go env GOPATH)/bin"

BIN_DIR      := bin
GO           := go
PROTOC       := /opt/homebrew/bin/protoc
COMPOSE_FILE := ./postgres/docker-compose.yml

.PHONY: all proto protoauth protogateway build-auth build-gateway run-auth run-gateway docker-up docker-down clean

## ====================== PROTO ==========================

proto: protoauth protogateway

protoauth:
	@export PATH="$$PATH:$$($(GO) env GOPATH)/bin"
	$(PROTOC) \
		--proto_path=auth-service/proto \
		--go_out=paths=source_relative:auth-service/proto \
		--go-grpc_out=paths=source_relative:auth-service/proto \
		auth-service/proto/auth.proto
	@echo "✅"

protogateway:
	@export PATH="$$PATH:$$($(GO) env GOPATH)/bin"
	$(PROTOC) \
		--proto_path=gateway/proto-clients/auth \
		--go_out=paths=source_relative:gateway/proto-clients/auth \
		--go-grpc_out=paths=source_relative:gateway/proto-clients/auth \
		gateway/proto-clients/auth/auth.proto
	@echo "✅"

## ====================== BUILD ==========================

build-auth:
	$(GO) build -o $(BIN_DIR)/auth-service ./auth-service/cmd/main.go
	@echo "✅"

build-gateway:
	$(GO) build -o $(BIN_DIR)/gateway ./gateway/cmd/main.go
	@echo "✅"

## ====================== RUN ==========================

run-auth: build-auth
	./$(BIN_DIR)/auth-service

run-gateway: build-gateway
	./$(BIN_DIR)/gateway

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
