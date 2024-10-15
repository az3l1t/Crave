FROM golang:1.22.5 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o order-service ./cmd/main.go

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/order-service .

COPY ./cmd /app/configs

EXPOSE 8081

CMD ["./order-service"]
