# Use a multi-stage build
FROM golang:1.23.2-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o backend ./main.go

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/backend .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./backend"]
