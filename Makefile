# Makefile for gio-shadcn project

# Default target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  demo                - Run interactive 37-component gallery demo"
	@echo "  build-demo          - Build demo app binary into bin/demo-app"
	@echo "  fmt FILE=<file>     - Format a specific Go file"
	@echo "  lint FILE=<file>    - Lint a specific Go file"
	@echo "  check FILE=<file>   - Run both fmt and lint on a file"
	@echo "  check-all           - Run fmt and lint on all Go files"
	@echo ""
	@echo "Usage examples:"
	@echo "  make demo"
	@echo "  make build-demo"
	@echo "  make fmt FILE=main.go"
	@echo "  make lint FILE=main.go"
	@echo "  make check FILE=main.go"
	@echo "  make check-all"

# Run interactive component gallery demo
.PHONY: demo run-demo
demo: run-demo
run-demo:
	@echo "Launching gio-shadcn component gallery demo..."
	go run ./demo/cmd

# Build component gallery demo binary into bin/
.PHONY: build-demo
build-demo:
	@echo "Building demo app binary into bin/demo-app..."
	@mkdir -p bin
	go build -o bin/demo-app ./demo/cmd

# Format a specific file
.PHONY: fmt
fmt:
	@if [ -z "$(FILE)" ]; then \
		echo "Error: FILE parameter is required. Usage: make fmt FILE=<filename>"; \
		exit 1; \
	fi
	@echo "Formatting $(FILE)..."
	go fmt $(FILE)

# Lint a specific file
.PHONY: lint
lint:
	@if [ -z "$(FILE)" ]; then \
		echo "Error: FILE parameter is required. Usage: make lint FILE=<filename>"; \
		exit 1; \
	fi
	@echo "Linting $(FILE)..."
	golangci-lint run $(FILE)

# Run both fmt and lint on a file
.PHONY: check
check: fmt lint
	@echo "Formatting and linting complete for $(FILE)"

# Run fmt and lint on all Go files
.PHONY: check-all
check-all:
	@echo "Formatting all Go files..."
	go fmt ./...
	@echo "Linting all Go files..."
	golangci-lint run ./...
	@echo "Formatting and linting complete for all files"
