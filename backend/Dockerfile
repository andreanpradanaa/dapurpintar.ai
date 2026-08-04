# --- build stage ---
FROM golang:1.26-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/server ./cmd/server

# --- runtime stage ---
FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata wget

COPY --from=builder /out/server /app/server
COPY data /app/data
COPY migrations /app/migrations

ENV PORT=8080
EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -q --spider http://localhost:${PORT}/api/v1/health || exit 1

CMD ["/app/server"]
