# Start from the official Go image for building
FROM golang:1.22 AS builder

# Set working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app (statically linked)
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Create a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./server"]
