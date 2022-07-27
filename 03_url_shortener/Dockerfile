FROM golang:1.18.4-alpine3.16

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o url_shortener main.go

EXPOSE 8080

CMD ["./url_shortener"]