# load environment variables
include .env
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' .env )
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))

pgup:
	@docker run -d \
	--name accounts_db \
	-e POSTGRES_PASSWORD=$(DB_PASS) \
	-e POSTGRES_DB=$(DB_NAME) \
	-p $(DB_PORT):$(DB_PORT) \
  postgres

pgdown:
	@docker stop accounts_db

composeup:
	@docker-compose up --build -d

composedown:
	@docker-compose down

migrateup:
	@migrate \
	-path ./infra/db/migrations \
	-database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)"  \
	-verbose up

migratedown:
	@migrate \
	-path ./infra/db/migrations \
	-database "postgresql://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)"  \
	-verbose down

migrategen:
	@migrate create -ext sql -dir infra/db/migrations create-accounts-table

test:
	@go test -v -cover ./...

install:
	@go get ./...

server:
	@go run cmd/api/main.go

build:
	@rm -rf ./bin/app && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./bin/app -ldflags="-s -w" ./cmd/api/main.go

.PHONY: pgdown pgup test install server build composeup composedown migrateup migratedown migrategen