FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go get github.com/silenceper/gowatch

ENTRYPOINT gowatch -o ./bin/main -p ./cmd