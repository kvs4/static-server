FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /static-server ./cmd/main.go

#-------------------------------------
FROM alpine:latest

COPY --from=builder /static-server /static-server

COPY --from=builder /app/.env .env

COPY --from=builder /app/static /static

EXPOSE 8080

CMD ["/static-server"]