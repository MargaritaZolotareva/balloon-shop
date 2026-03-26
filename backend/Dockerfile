FROM golang:1.24-alpine as go-build

RUN apk add --no-cache git curl tar

WORKDIR /app
COPY . .

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.0/migrate.linux-amd64.tar.gz \
    | tar xvz \
    && mv migrate /usr/local/bin/migrate
RUN go mod tidy && go build -o balloon-shop cmd/main.go

RUN go build -o fill_slugs ./scripts/slugs/fill_slugs.go
RUN go build -o generate_sitemap ./scripts/sitemap/generate_sitemap.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /build

COPY --from=go-build /app/balloon-shop .
COPY --from=go-build /app/fill_slugs .
COPY --from=go-build /app/generate_sitemap .
COPY --from=go-build /usr/local/bin/migrate /usr/local/bin/migrate