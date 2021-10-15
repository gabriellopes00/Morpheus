migrategen:
	@migrate create -ext sql -dir ./framework/db/migrations $(migration_name)

test:
	@go test -v -cover ./...

install:
	@go get ./...

server:
	@go run cmd/server.go

build:
	@rm -rf ./app && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./app -ldflags="-s -w" ./cmd/server.go

.PHONY: migrategen test install server build
