# Build automation for the app.  Usage: make <target>
APP      := app
BIN_DIR  := bin
PKG      := ./cmd/app
VERSION  := $(shell git describe --tags --always 2>/dev/null || echo dev)
LDFLAGS  := -ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: all run build test fmt vet clean help

all: fmt vet test build ## Run fmt, vet, test and build

run: ## Run the application
	go run $(PKG)

build: ## Compile binary into bin/
	@mkdir -p $(BIN_DIR)
	go build $(LDFLAGS) -o $(BIN_DIR)/$(APP) $(PKG)

test: ## Run tests with coverage
	go test -v -cover ./...

fmt: ## Format source code
	go fmt ./...

vet: ## Static analysis
	go vet ./...

clean: ## Remove build output
	rm -rf $(BIN_DIR)

help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN{FS=":.*?## "}{printf "  %-8s %s\n", $$1, $$2}'
