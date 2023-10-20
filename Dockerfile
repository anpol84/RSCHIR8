
# Используем образ golang:latest как базовый
FROM golang:latest

# Устанавливаем переменные среды
ENV PORT=8080
ENV NAME=name

# Создаем директорию внутри контейнера, куда будет копироваться код приложения
WORKDIR /app

# Копируем файлы go.mod и go.sum в директорию /app
COPY go.mod go.sum ./

# Запускаем команду go mod download для загрузки зависимостей
RUN go mod download

# Копируем все файлы проекта в директорию /app
COPY . .



# Указываем команду, которая будет исполняться при запуске контейнера
CMD go run ./cmd/main.go --port=${PORT} --name=${NAME}