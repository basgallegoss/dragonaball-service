# ── Builder (Go 1.23) ───────────────────────────────────────────────
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código y compila
COPY . .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags="-s -w" -o dragonball cmd/server/main.go

# ── Runtime (Alpine con CA certs) ──────────────────────────────────
FROM alpine:latest
# Instala los certificados raíz para TLS
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/dragonball .

EXPOSE 8080
ENTRYPOINT ["./dragonball"]
