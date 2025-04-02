generate:
	mkdir -p internal/generated/schema
	oapi-codegen -package rider -generate chi-server,types,spec -o internal/generated/schema/rider.gen.go api/schema.yaml

install-deps:
	go install github.com/jackc/tern/v2@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

migrate:
	cd internal/db/migrations && tern migrate

up:
	docker-compose up -f build/docker-compose.yml up -d
	DATABASE_URL="postgresq://postgres:postgres@localhost:6432/postgres" \
	go run cmd/service/main.go

stop:
	docker compose -f build/docker-compose.yml stop
	