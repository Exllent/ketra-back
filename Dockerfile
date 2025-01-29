# Stage 1: Build the application
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ketra-back .

# Stage 2: Create the final minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/ketra-back .

# Copy .env file (optional, if needed for development)
COPY .env .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./ketra-back"]
