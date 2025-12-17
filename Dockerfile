FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin /app/bin
CMD ["/app/bin"]
