# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o cloudcli ./cmd/cli

# Final stage
FROM scratch
COPY --from=builder /app/server /server
COPY --from=builder /app/cloudcli /cloudcli

EXPOSE 8080
ENTRYPOINT ["/server"]

