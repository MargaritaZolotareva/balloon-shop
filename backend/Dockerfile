FROM node:20-alpine AS frontend-build

WORKDIR /app

ARG VUE_APP_URL
ARG VUE_APP_API_URL

COPY frontend/package*.json ./
RUN npm install
COPY frontend .
RUN VUE_APP_URL=$VUE_APP_URL VUE_APP_API_URL=$VUE_APP_API_URL npm run build

FROM golang:1.24-alpine as builder

RUN apk add --no-cache git curl tar

WORKDIR /app

COPY backend .

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.0/migrate.linux-amd64.tar.gz \
    | tar xvz \
    && mv migrate /usr/local/bin/migrate
RUN go mod tidy && go build -o balloon-shop cmd/main.go

RUN go build -o fill_slugs ./scripts/fill_slugs.go
RUN go build -o generate_sitemap ./scripts/generate_sitemap.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /build

COPY --from=frontend-build /app/dist ./dist
COPY --from=builder /app/balloon-shop .
COPY --from=builder /app/fill_slugs .
COPY --from=builder /app/generate_sitemap .
COPY --from=builder /app/photos ./photos
COPY --from=builder /app/infrastructure/db/migrations ./migrations

COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

ENTRYPOINT ["/bin/sh", "-c", "\
    echo 'Running migrations'; \
    migrate -path ./migrations -database \"$DATABASE_URL\" up; \
    echo 'Running fill_slugs'; \
    ./fill_slugs; \
    echo 'Running generate_sitemap'; \
    ./generate_sitemap; \
    echo 'Starting backend'; \
    ./balloon-shop \
"]