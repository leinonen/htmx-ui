# Build stage
FROM golang:1.24.1 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy

# 🔧 Add these environment variables to force static binary
ENV CGO_ENABLED=0 GOOS=linux

RUN go build -o html_server main.go

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/html_server .
COPY templates ./templates

EXPOSE 8080

CMD ["./html_server"]
