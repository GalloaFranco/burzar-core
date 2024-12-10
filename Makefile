APP_NAME := main
GOOS := linux
GOARCH := amd64
ENV ?= local

.PHONY: clean build run-local run-proxy format

# Remove the binary if it exists
clean:
	@if [ -f $(APP_NAME) ]; then \
		echo "Removing existing binary..."; \
		rm -f $(APP_NAME); \
	fi

# Compile the application for AWS Lambda
build: clean
	@echo "Building application..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_NAME)
	@echo "Build completed: $(APP_NAME)"

# Run the REST application locally from main.go
run-local: clean
	@echo "Running REST application locally with environment..."
	ENV=$(ENV) go run main.go

# Run SAM CLI in Lambda Proxy mode
run-proxy: build
	@echo "Running application in Lambda Proxy mode..."
	SAM_CLI_ENV=$(ENV) sam local start-api --env-vars env.json


# Format code with goimports and gofmt
format:
	@echo "Running goimports...."
	goimports -w .
	@echo "Running gofmt..."
	gofmt -e -w .
	@echo "Code formatting completed."
