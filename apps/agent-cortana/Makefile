.PHONY: build run test clean docker-build docker-run help

# Variables
APP_NAME=agent-cortana
BINARY_NAME=agent-cortana
DOCKER_IMAGE=agent-cortana:latest

# Default target
help:
	@echo "Agent Cortana - Makefile commands:"
	@echo "  make build        - Build the application binary"
	@echo "  make run          - Run the application locally"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	go build -o $(BINARY_NAME) ./cmd/server

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	go run ./cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	go clean

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) -f Dockerfile ../..

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker-compose up
