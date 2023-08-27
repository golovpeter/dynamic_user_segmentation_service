FROM golang:1.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

# Аргумент для передачи пути к файлу
ARG MAIN_FILE_PATH
ARG EXEC_FILE_NAME

# Собираем приложение внутри контейнера
RUN go build -o $EXEC_FILE_NAME $MAIN_FILE_PATH
