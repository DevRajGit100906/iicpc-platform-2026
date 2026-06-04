# Stage 1: Build the binary
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Copy the SUT code
COPY main.go .

# Initialize a temporary go module and build
RUN go mod init dummy-sut && \
    go build -o matching-engine main.go

# Stage 2: Create the minimal production image
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/matching-engine .

# Expose the port the SUT listens on
EXPOSE 8080

# Run the binary
CMD ["./matching-engine"]