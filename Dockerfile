# ---------- Stage 1: Build the Go binary ----------
FROM golang:1.23.0 AS builder

# Set Go environment for static binary
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./

# Download all dependencies and cache them
RUN go mod tidy

# Copy the rest of the code
COPY . .

# Build the Go binary
RUN go build -o /inventory-service ./cmd/grpc

# ---------- Stage 2: Create a minimal runtime image ----------
FROM alpine:3.18

# Set working directory in container
WORKDIR /app

# Copy compiled binary from builder
COPY --from=builder /inventory-service .

# Expose gRPC port
EXPOSE 50051

# Run the service
CMD ["./inventory-service"]