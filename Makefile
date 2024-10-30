BINARY_NAME=subnet-calc
VERSION?=1.0.0
GOPATH=$(shell go env GOPATH)

.PHONY: build clean test coverage lint help

build: ## Build the binary
	go build -o $(BINARY_NAME) -v

test: ## Run tests
	go test -v ./...

coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint: ## Run linter
	$(GOPATH)/bin/golangci-lint run

clean: ## Clean build files
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help