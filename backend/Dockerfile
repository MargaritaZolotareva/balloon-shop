FROM golang:1.24-alpine as builder

RUN apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o balloon-shop cmd/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /build

COPY --from=builder /app/balloon-shop .