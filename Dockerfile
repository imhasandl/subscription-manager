FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o subscription-manager .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/subscription-manager .
COPY .env .

EXPOSE 8080

CMD ["./subscription-manager"]