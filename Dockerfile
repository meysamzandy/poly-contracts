# Dockerfile for protobuf tooling
FROM golang:1.22-alpine AS builder

# Install buf and protoc
RUN apk add --no-cache protoc buf

# Set working directory
WORKDIR /workspace
COPY . /workspace

# Generate code
RUN buf generate

# Default command
CMD ["protoc"]
