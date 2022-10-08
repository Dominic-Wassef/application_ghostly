BINARY_NAME=ghostlyApp

build:
	@go mod vendor
	@echo "Building Ghostly..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Ghostly built!"

run: build
	@echo "Starting Ghostly..."
	@./tmp/${BINARY_NAME} &
	@echo "Ghostly started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Ghostly..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Ghostly!"

restart: stop start