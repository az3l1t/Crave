FROM golang:1.22.5 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o auth-service ./cmd/main.go

FROM debian:buster

WORKDIR /app

COPY --from=builder /app/auth-service .

COPY ./configs /app/configs

EXPOSE 8080

CMD ["/.auth-service"]