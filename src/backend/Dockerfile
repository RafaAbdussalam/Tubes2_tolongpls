# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api/

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /api /app/api
COPY ./data/alchemy.csv /app/data/
COPY ./data/alchemy.db /app/data/
EXPOSE 8080
CMD ["/app/api"]