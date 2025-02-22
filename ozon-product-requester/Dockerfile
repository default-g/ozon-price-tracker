FROM golang:1.21-alpine AS builder

WORKDIR /app

# Установка необходимых системных пакетов
RUN apk add --no-cache \
    chromium \
    chromium-chromedriver \
    ca-certificates \
    tzdata

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
COPY --from=builder /usr/bin/chromium-browser /usr/bin/
COPY --from=builder /usr/bin/chromedriver /usr/bin/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Установка необходимых зависимостей
RUN apk add --no-cache \
    chromium \
    chromium-chromedriver \
    ca-certificates \
    tzdata

# Установка переменных окружения для Chrome
ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/bin/chromium-browser \
    CHROMEDRIVER_PATH=/usr/bin/chromedriver

# Запуск приложения
CMD ["./product-info"] 