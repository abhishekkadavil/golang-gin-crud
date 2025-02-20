# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy only the Go source files and go.mod/go.sum
COPY . .

# Download and install dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Stage 2: Create the final image with only the binary
FROM alpine:latest

WORKDIR /app    

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file from the build context into the container (if it exists)
COPY .env .env

# Expose the port the app runs on
EXPOSE 3000

# Run the executable
CMD ["./main"]