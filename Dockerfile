# Stage 1: Build the Go application
FROM golang:1.17-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 ensures a statically linked binary
# -ldflags="-s -w" strips debugging information and symbols to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o timecron main.go

# Verify that the executable (expected as 'main' based on logs) exists in the builder stage
RUN ls -l /app/main && echo "Builder: /app/main (the actual executable) verified."

# Stage 2: Create the runtime image
FROM alpine:latest

# Install timezone data and bash
RUN apk --no-cache add tzdata bash

WORKDIR /app

# Copy the compiled binary (named 'main' in builder) from the builder stage
# and name it 'timecron' in the runtime stage.
COPY --from=builder /app/main ./timecron

# Verify the copied file (now named timecron) and ensure it's executable
RUN ls -l /app/timecron && chmod +x /app/timecron

# Copy the configuration file
# The application will create a default config.json if it's not found,
# but it's good practice to include a base version or ensure it can be mounted.
COPY config.json ./config.json

# Copy the scripts directory if your tasks in config.json depend on external scripts
# Adjust if your scripts are located elsewhere or not needed.
COPY script ./script

# Create a logs directory (can be mounted as a volume)
RUN mkdir -p /app/logs

# Expose the port the application runs on (default is 3005, can be overridden in config.json)
EXPOSE 3005

# Command to run the application
CMD sh -c "echo 'Listing /app contents:'; ls -la /app; echo 'Attempting to execute /app/timecron:'; /app/timecron ; echo \"Timecron application exited with status $?. Sleeping indefinitely.\" ; sleep infinity" 