FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates tzdata git build-base

WORKDIR /src

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

ENV CGO_ENABLED=0
ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -v -trimpath -ldflags="-s -w" -o /out/app ./cmd



FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -H -u 10001 joshb

WORKDIR /home/app

COPY --from=builder /out/app /usr/local/bin/app

USER joshb

ENV PORT=8080
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
    CMD wget -q0- "http://127.0.0.1:${PORT}/health" >/dev/null || exit 1

ENTRYPOINT ["/usr/local/bin/app"]