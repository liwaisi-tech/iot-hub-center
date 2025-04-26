# syntax=docker/dockerfile:1

# --- Build Stage ---
FROM golang:1.24.2 AS builder
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o mqtt-app ./cmd/app/main.go

# --- Final Stage ---
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/mqtt-app ./mqtt-app

# If you have .env or config files, uncomment the next line
# COPY .env ./

# Expose any ports your app uses (example: 8080)
# EXPOSE 8080

ENTRYPOINT ["./mqtt-app"]
