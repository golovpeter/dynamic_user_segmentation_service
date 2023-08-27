FROM golang:1.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# Аргумент для передачи пути к файлу
ARG APP_PATH
ARG FILE_NAME

# Собираем приложение внутри контейнера
RUN go build -o $FILE_NAME $APP_PATH

EXPOSE 8080

# Команда по умолчанию для запуска приложения
#CMD ["./$FILE_NAME"]