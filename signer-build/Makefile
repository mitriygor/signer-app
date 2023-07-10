SIGNER_API=signerApi
BROKER_SERVICE=brokerService
LOGGER_SERVICE=loggerService
LISTENER_SERVICE=listenerService

# Builds the Signer API binary as a Linux executable
signer_api:
	@echo "Building Signer API binary..."
	cd ../signer-api && env GOOS=linux CGO_ENABLED=0 go build -o ${SIGNER_API} ./cmd/api
	@echo "Signer API is built..."

## build_broker: builds the broker binary as a linux executable
broker_service:
	@echo "Building broker service..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_SERVICE} ./cmd/api
	@echo "Done!"

# Builds the Logger binary as an executable
logger_service:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_SERVICE} ./cmd/api
	@echo "Logger service is built"

## listener_service: builds the listener service as a linux executable
listener_service:
	@echo "Building listener binary..."
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_SERVICE} .
	@echo "Done!"

# Stops docker-compose, — if it runs, — and builds all parts, and starts Docker compose
build_up: signer_api broker_service logger_service listener_service
	@echo "Stopping Docker images..."
	docker-compose down
	@echo "Building, — when required, — and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images are built and started"


# Starts all containers in the background
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images are started"


# Stops docker compose
down:
	@echo "Stopping Socker compose..."
	docker-compose down
	@echo "Docker images are down..."










