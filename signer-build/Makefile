BROKER_SERVICE=brokerService
LISTENER_SERVICE=listenerService
KEY_KEEPER_SERVICE=keyKeeperService
SIGNER_API=signerApi
LOGGER_SERVICE=loggerService


# build_broker: builds the Broker binary as a Linux executable
broker_service:
	@echo "Building broker service..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_SERVICE} ./cmd/api
	@echo "Broker Service is built..."


# listener_service: builds the Listener service as a Linux executable
listener_service:
	@echo "Building listener binary..."
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_SERVICE} ./cmd/api
	@echo "Listener Service is built..."


# key_keeper_service: builds the Signer API binary as a Linux executable
key_keeper_service:
	@echo "Building Key Keeper Service binary..."
	cd ../key-keeper-service && env GOOS=linux CGO_ENABLED=0 go build -o ${KEY_KEEPER_SERVICE} ./cmd/api
	@echo "Key Keeper Service is built..."


# signer_api: builds the Signer API binary as a Linux executable
signer_api:
	@echo "Building Signer API binary..."
	cd ../signer-api && env GOOS=linux CGO_ENABLED=0 go build -o ${SIGNER_API} ./cmd/api
	@echo "Signer API is built..."


# logger_service: builds the Logger binary as an executable
logger_service:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_SERVICE} ./cmd/api
	@echo "Logger Service is built"


# Stops docker-compose, — if it runs, — and builds all parts, and starts Docker compose
build_up: broker_service listener_service key_keeper_service signer_api logger_service
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










