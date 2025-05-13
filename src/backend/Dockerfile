# Stage 1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /scraper ./cmd/scraper/main.go

# Stage 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /api /app/api
COPY --from=builder /scraper /app/scraper
COPY --from=builder /app/data /app/data
EXPOSE 8080