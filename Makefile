install-lint-tool:
	@echo "  >  Installing lint tool..."
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.50.1

lint: install-lint-tool
	@echo "  >  Linting..."
	@bin/golangci-lint run

lint-fix: install-lint-tool
	@echo "  >  Linting and fixing..."
	@bin/golangci-lint run --fix

test:
	@echo "  >  Running tests..."
	@go test ./...

.PHONY:	build
build:
	@echo "  >  Building project binary..."
	@go build

.PHONY:	start
## start the invasion with 2 aliens
start: build
	@./alien-invasion  2

.PHONY: go-mod-download
go-mod-download:
	@echo "  >  Downloading dependencies..."
	@go mod download
