# Stage 1: Build the Go application
FROM golang:alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and change to the app directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main cmd/web/main.go

# Stage 2: Run the Go application
FROM alpine:latest

# Install make and other dependencies
RUN apk add --no-cache make bash

# Create and change to the app directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the config files
COPY config.development.json .
COPY config.production.json .

# Expose the port the app runs on
EXPOSE 5000

# Run the application
CMD ["./main"]