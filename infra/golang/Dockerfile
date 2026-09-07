FROM golang:latest

# Установка Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Установка Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Установка Air
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Открываем порт для дебага
EXPOSE 2345 8000

CMD ["sleep", "infinity"]
