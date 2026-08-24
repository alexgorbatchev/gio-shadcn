# Justfile for gio-shadcn

default:
    @just --list

# Run interactive 37-component gallery demo
demo:
    @echo "Launching gio-shadcn component gallery demo..."
    go run ./demo/cmd

# Alias for demo
run: demo

# Build component gallery demo binary into bin/demo-app
build-demo:
    @echo "Building demo app binary into bin/demo-app..."
    @mkdir -p bin
    go build -o bin/demo-app ./demo/cmd

# Run unit tests across all packages
test:
    go test ./...

# Run static analysis
vet:
    go vet ./...

# Format a specific file
fmt file:
    go fmt {{file}}

# Format all Go files
fmt-all:
    go fmt ./...

# Lint a specific file
lint file:
    golangci-lint run {{file}}

# Lint all Go files
lint-all:
    golangci-lint run ./...

# Run both fmt and lint on all Go files
check-all: fmt-all lint-all
