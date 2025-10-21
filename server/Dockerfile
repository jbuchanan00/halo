# ---- builder ----
FROM golang:1.25-bookworm AS builder
WORKDIR /src

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .
COPY /cmd/adhoc/processGeoData/cityStateToLatLong.json /cityStateToLatLong.json

# Build a static linux binary (musl-friendly) and put it in /out
# Adjust CGO_ENABLED/GOARCH if you need a different arch
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /out/halo ./cmd

# ---- runtime ----
FROM alpine:3.19 AS runtime

# create non-root user and ensure home exists
RUN adduser -D -u 1000 appuser

WORKDIR /home/appuser

# copy binary, set ownership to appuser, and make sure it's executable
COPY --from=builder --chown=appuser:appuser /out/halo ./halo
COPY --from=builder /cityStateToLatLong.json /app/data/cityStateToLatLong.json

RUN chmod +x ./halo

ENV PORT=8080
EXPOSE 8080

USER appuser

HEALTHCHECK --interval=30s --timeout=3s CMD curl -f http://127.0.0.1:8080/api/health_checks/ready || exit 1

ENTRYPOINT ["./halo"]
