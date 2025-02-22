FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копирование файлов проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o product-info ./cmd/product-info

# Финальный образ
FROM alpine:latest

WORKDIR /app

# Копирование необходимых файлов из builder
COPY --from=builder /app/product-info .

# Установка необходимых зависимостей
RUN apk add --no-cache \
    chromium \
    chromium-chromedriver \
    ca-certificates \
    tzdata

# Запуск приложения
CMD ["./product-info"]