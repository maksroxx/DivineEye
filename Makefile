BIN_DIR      := bin
GO           := go
PROTOC       := /opt/homebrew/bin/protoc

.PHONY: protoauth

protoauth:
	@export PATH="$$PATH:$$($(GO) env GOPATH)/bin"
	$(PROTOC) \
		--proto_path=auth-service/proto \
		--go_out=paths=source_relative:auth-service/proto \
		--go-grpc_out=paths=source_relative:auth-service/proto \
		auth-service/proto/auth.proto
