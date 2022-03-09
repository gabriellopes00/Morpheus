FROM golang:1.17.0-alpine3.14 AS builder

WORKDIR /go/src/morpheus/accounts

COPY go.* ./
COPY Makefile .

RUN apk add --no-cache make
RUN make install 

COPY .env .
COPY . .
RUN make build

FROM alpine:3.14

WORKDIR /usr/src/morpheus/accounts

COPY --from=builder /go/src/morpheus/accounts/app .
COPY .env .
COPY ./pkg/db/migrations/*.sql ./pkg/db/migrations/

EXPOSE 3792
ENTRYPOINT /usr/src/morpheus/accounts/app
