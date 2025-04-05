generate:
	mkdir -p internal/generated/schema
	mkdir -p internal/generated/proto
	oapi-codegen -package rider -generate chi-server,types,spec -o internal/generated/schema/rider.gen.go api/schema.yaml
	protoc --go_out=./internal/generated/proto \
	--go-grpc_out=./internal/generated/proto \
	api/proto/service.proto

install-deps:
	go install github.com/jackc/tern/v2@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	brew install protobuf

migrate:
	cd internal/db/migrations && tern migrate

up:
	docker-compose up -f build/docker-compose.yml --project-name rider up -d
	DATABASE_URL="postgresq://postgres:postgres@localhost:6432/postgres" \
	GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info \
	go run cmd/service/main.go

stop:
	docker compose -f build/docker-compose.yml stop
