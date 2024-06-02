SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: tidy proto build run

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR)/product-backend -v ./cmd/server

run: ## Start application
	$(GOCMD) run ./cmd/server

server: ## Start server application
	$(GOCMD) run ./cmd/server

test: ## Run tests
	$(GOCMD) test ./... -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

deps: ## Install dependencies
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

tidy: ## init go modules
	$(GOCMD) mod tidy && $(GOCMD) mod vendor

proto:
	cd internal/delivery/grpc/proto && echo "init grpc" && protoc *.proto --go_out=../../../../. --go-grpc_out=../../../../. && \
	echo "init grpc-gateway" && protoc --grpc-gateway_out=logtostderr=true:../../../../. *.proto && \
	echo "init swagger" && protoc -I . --openapiv2_out ../../swagger *.proto && \
	echo "finish make proto"

