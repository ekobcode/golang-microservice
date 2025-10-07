# Dockerfile (Multi-stage build for Golang microservice)
# ---------- STAGE 1: Build ----------
FROM golang:1.22.5-alpine AS builder


# Enable Go modules and set environment variables
ENV CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64


WORKDIR /app


# Install dependencies
RUN apk add --no-cache git


# Copy go.mod and go.sum first (leverage Docker cache)
COPY go.mod go.sum ./
RUN go mod download && go mod verify


# Copy the rest of the source code
COPY . .


# Build the binary
RUN go build -o server ./cmd/server/main.go


# ---------- STAGE 2: Run ----------
FROM alpine:3.20


# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup


WORKDIR /app


# Copy compiled binary and environment file
COPY --from=builder /app/server ./server
COPY .env.example .env


# Expose application port
EXPOSE 8080


# Set permissions and run as non-root
RUN chown -R appuser:appgroup /app
USER appuser


# Command to run the app
ENTRYPOINT ["./server"]