FROM golang:1.17.0-alpine3.14 AS builder

WORKDIR /go/src/morpheus/tickets

RUN adduser -D morpheus && chown -R morpheus /go/src/morpheus/tickets
USER morpheus

COPY go.* ./
COPY Makefile .

RUN apk add --no-cache make
RUN make install 

COPY . .
RUN make build

FROM alpine:3.14

WORKDIR /usr/src/morpheus/tickets

COPY --from=builder /go/src/morpheus/tickets/app .
COPY .env .
COPY ./framework/db/migrations/*.sql ./framework/db/migrations/

ENTRYPOINT /usr/src/morpheus/tickets/app
