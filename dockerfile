# Step 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o crudapi main.go

# Step 2: Run the Go binary
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built Go binary from the builder container
COPY --from=builder /app/crudapi .

# Expose the port that the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./crudapi"]
