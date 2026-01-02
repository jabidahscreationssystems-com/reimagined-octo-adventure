# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy the app source
COPY apps/agent-cortana ./apps/agent-cortana

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o agent-cortana ./apps/agent-cortana/cmd/server

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/agent-cortana .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./agent-cortana"]
