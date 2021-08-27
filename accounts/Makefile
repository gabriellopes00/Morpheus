pgup:
	@docker run -d\
	--name accounts_db \
	-e POSTGRES_PASSWORD=accounts_postgres \
	-e POSTGRES_DB=accounts \
	-v accounts_db:/var/lib/postgresql/data/ \
	-p 5432:5432 \
  postgres:13.4-alpine3.14

pgdown:
	@docker stop accounts_db

composeup:
	@docker-compose up --build -d

composedown:
	@docker-compose down

test:
	@go test -v -cover ./...

install:
	@go get ./...

server:
	@go run cmd/api/main.go

build:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./bin/app -ldflags="-s -w" ./cmd/api/main.go

.PHONY: pgdown pgup test install server build composeup composedown