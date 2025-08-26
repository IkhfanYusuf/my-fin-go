# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o app

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
