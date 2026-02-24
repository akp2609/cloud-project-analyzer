FROM golang:1.25.5-alpine AS builder

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
COPY vendor/ ./vendor/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod=vendor -o logs-ingestor ./cmd/logs-ingestor

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/logs-ingestor /logs-ingestor
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/logs-ingestor"]