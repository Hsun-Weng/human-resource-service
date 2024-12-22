APP_NAME=human-resource-service
IMAGE_NAME=human-resource-service
TAG=$(shell git rev-parse --short HEAD)
VERSION=1.0.0
SRC_DIR=.

.PHONY: all
all: docker-build docker-compose-up

.PHONY: docker-build
docker-build:
	@echo "Building Docker image $(IMAGE_NAME):$(TAG)..."
	docker build -t $(IMAGE_NAME):$(TAG) .

.PHONY: docker-push
docker-push:
	@echo "Pushing Docker image $(IMAGE_NAME):$(TAG)..."
	docker push $(IMAGE_NAME):$(TAG)

.PHONY: docker-compose-up
docker-compose-up:
	@echo "Starting application using Docker Compose..."
	docker compose up -d

.PHONY: docker-compose-down
docker-compose-down:
	@echo "Stopping application..."
	docker compose down

.PHONY: deploy
deploy: docker-build docker-push docker-compose-up
	@echo "Deploying application to production..."

.PHONY: test
test:
	@echo "Running unit tests..."
	go test ./...

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
	docker rmi $(IMAGE_NAME):$(TAG) || true

.PHONY: help
help:
	@echo "Makefile Commands:"
	@echo "  make docker-build      Build the Docker image"
	@echo "  make docker-push       Push Docker image to registry"
	@echo "  make docker-compose-up Start the application with Docker Compose"
	@echo "  make docker-compose-down Stop the application"
	@echo "  make deploy            Build, Push, and Deploy to production"
	@echo "  make test              Run Go tests"
	@echo "  make clean             Clean up build files and Docker images"
