FROM golang:1.22.5 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api-gateway ./cmd/main.go

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/api-gateway .

COPY ./cmd /app/configs

EXPOSE 8080

CMD ["./api-gateway"]
