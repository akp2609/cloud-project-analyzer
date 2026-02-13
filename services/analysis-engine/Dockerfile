FROM golang:1.25.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy go.mod, go.sum, and vendor folder
COPY go.mod go.sum ./
COPY vendor/ ./vendor/
COPY . .

# Build using vendored dependencies
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod=vendor -o analysis-engine ./cmd/analysis-engine

# Final stage: minimal distroless image
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/analysis-engine /analysis-engine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/analysis-engine"]
