ENTRYPOINT_PATH=./cmd/shell

GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO_COVERAGE_OUTPUT=coverage.out
GO_COVERAGE_HTML_OUTPUT=coverage.html

LINUX_BUILD_PATH=./dist/linux
WINDOWS_BUILD_PATH=./dist/windows
MAC_BUILD_PATH=./dist/mac

BINARY_NAME=ohmygosh
BINARY_LINUX=$(BINARY_NAME)-linux
BINARY_WINDOWS=$(BINARY_NAME)-windows
BINARY_MAC=$(BINARY_NAME)-mac

MSG_BUILD_COMPLETED := "Build completed."

.PHONY: all
all: deps fmt test build

.PHONY: build
build: deps
	@echo "Building for current platform..."
	$(GOBUILD) -o $(BINARY_NAME) $(ENTRYPOINT_PATH)/main.go
	@echo $(MSG_BUILD_COMPLETED)

.PHONY: build-linux
build-linux: deps
	@echo "Building for Linux..."
	mkdir -p $(LINUX_BUILD_PATH)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(LINUX_BUILD_PATH)/$(BINARY_LINUX) $(ENTRYPOINT_PATH)/main.go
	@echo $(MSG_BUILD_COMPLETED)

.PHONY: build-windows
build-windows: deps
	@echo "Building for Windows..."
	mkdir -p $(WINDOWS_BUILD_PATH)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(WINDOWS_BUILD_PATH)/$(BINARY_WINDOWS) $(ENTRYPOINT_PATH)/main.go
	@echo $(MSG_BUILD_COMPLETED)

.PHONY: build-mac
build-mac: deps
	@echo "Building for Mac..."
	mkdir -p $(MAC_BUILD_PATH)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(MAC_BUILD_PATH)/$(BINARY_MAC) $(ENTRYPOINT_PATH)/main.go
	@echo $(MSG_BUILD_COMPLETED)

.PHONY: build-all
build-all: build-linux build-windows build-mac

.PHONY: test
test:
	@echo "Running Tests..."
	$(GOTEST) -v ./...
	@echo "Completed."

.PHONY: coverage
coverage:
	@echo "Generating test coverage..."
	$(GOTEST) -coverprofile=$(GO_COVERAGE_OUTPUT) ./...
	$(GOCMD) tool cover -html=$(GO_COVERAGE_OUTPUT) -o $(GO_COVERAGE_HTML_OUTPUT)
	@echo "Coverage report generated."

.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf dist/*
	@echo "Cleanup completed."

.PHONY: fmt
fmt:
	@echo "Formatting project..."
	$(GOCMD) fmt ./...
	@echo "Formatting completed."

.PHONY: lint
lint:
	@echo "Linting..."
	@command -v golangci-lint >/dev/null 2>&1 || \
		(echo "Installing golangci-lint..." && \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin)
	golangci-lint run
	@echo "Linting completed."

.PHONY: run
run:
	$(GORUN) $(ENTRYPOINT_PATH)/main.go

.PHONY: deps
deps:
	$(GOCMD) mod tidy
	$(GOGET) -u ./...