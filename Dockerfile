# Use the official Golang image as the base image
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the application
RUN go build -o server

# Use a minimal Alpine image as the base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder image
COPY --from=builder /app/server .

# Expose the port that the server listens on
EXPOSE 8888

# Run the server
# Open http://huma-demo.traefik.me if you are using traefik reverse proxy with docker-compose
CMD ["./server"]

