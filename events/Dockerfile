FROM golang:1.17.0-alpine3.14 AS builder

WORKDIR /go/src/morpheus/events

COPY go.* ./
COPY Makefile .
COPY .env .

RUN apk add --no-cache make
RUN make install

COPY . .
RUN make build

FROM alpine:3.14

WORKDIR /usr/src/morpheus/events

COPY --from=builder /go/src/morpheus/events/app .
COPY .env .
COPY ./framework/db/migrations/*.sql ./framework/db/migrations/

EXPOSE 7765
ENTRYPOINT /usr/src/morpheus/events/app
