# Multi-stage build Dockerfile for Go
# --- Builder stage ---
FROM golang:alpine AS builder
WORKDIR /app

# Install git (required for go mod)
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o aprendiz-postgres ./cmd/aprendiz-postgres/

# --- Final stage ---
FROM alpine:latest
WORKDIR /app

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy the built binary from builder
COPY --from=builder /app/aprendiz-postgres ./aprendiz-postgres

# Change ownership and permissions
RUN chown appuser:appgroup ./aprendiz-postgres

# Expose port 8080 (optional, adjust as needed)
EXPOSE 8080

# Switch to non-root user
USER appuser

# Command to run the executable
CMD ["./aprendiz-postgres"]
