FROM golang:1.23.3

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go
COPY .env .env

EXPOSE 8080

CMD ["./main"]