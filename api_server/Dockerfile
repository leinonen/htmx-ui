# Build stage
FROM golang:1.24.1 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy

# 🔧 Add these environment variables to force static binary
ENV CGO_ENABLED=0 GOOS=linux

RUN go build -o api_server main.go

# Verify that the binary was created
RUN ls -l /app/api_server

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/api_server .

EXPOSE 8081

CMD ["./api_server"]
