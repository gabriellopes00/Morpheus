FROM golang:1.17.0-alpine3.14 AS builder

WORKDIR /go/src/morpheus/accounts

COPY . .

RUN apk update && apk upgrade && apk add --no-cache make
RUN make install && make build

FROM alpine:3.14

WORKDIR /usr/src/morpheus/accounts

COPY --from=builder /go/src/morpheus/accounts/bin/app .
COPY .env .

EXPOSE 7765
ENTRYPOINT /usr/src/morpheus/accounts/app