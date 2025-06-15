# Build stage
FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o your-app ./cmd/your-app

# Run stage
FROM alpine:latest

RUN adduser -D appuser
USER appuser

WORKDIR /home/appuser

COPY --from=builder /app/your-app .

EXPOSE 8080

ENTRYPOINT ["./halo"]
