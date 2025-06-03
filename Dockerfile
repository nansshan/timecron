# Stage 1: Build the Go application
FROM golang:1.17-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project source code
COPY . .

# Build the Go application and output to a specific path /app/timecron_executable
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/timecron_executable main.go

# List all contents of /app in the builder stage to see what was generated
RUN ls -la /app && echo "Builder: /app directory listed to check for timecron_executable."

# Stage 2: Create the runtime image
FROM alpine:latest

# Install timezone data and bash
RUN apk --no-cache add tzdata bash

WORKDIR /app

# Copy the compiled binary (assuming it's /app/timecron_executable in builder)
# and name it 'timecron' in the runtime stage.
COPY --from=builder /app/timecron_executable ./timecron

# Verify the copied file (now named timecron) and ensure it's executable
RUN ls -l /app/timecron && chmod +x /app/timecron

# Original COPY config.json is now replaced by the RUN echo command above for testing
# RUN echo '{ "name": "timecron-minimal", "username": "admin", "email": "xnkyn@qq.com", "password": "21232f297a57a5a743894a0e4a801fc3", "task": [] }' > /app/config.json

# Copy the configuration file from the build context (repository)
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