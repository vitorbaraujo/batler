GO ?= go
GOLANGCI_LINT := $(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint
GOLINT := $(GO) run golang.org/x/lint/golint
GOIMPORTS := $(GO) run golang.org/x/tools/cmd/goimports

SOURCES := $(shell \
	find . -not \( \( -name .git -o -name vendor \) -prune \) \
	-name *.go)

PACKAGES := $(shell $(GO) list ./...)

.PHONY: build
build:
	$(GO) build -o build/batler ./cmd/batler/

.PHONY: check
check:
	$(GO) test ./... -coverpkg=./... -coverprofile=coverage.out

.PHONY: fix
fix:
	$(GOIMPORTS) -w $(SOURCES)

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run
	$(GOLINT) -set_exit_status $(PACKAGES)

.PHONY: cover
cover: cover/text

.PHONY: cover/text
cover/text:
	$(GO) tool cover -func=coverage.out

.PHONY: cover/html
cover/html:
	$(GO) tool cover -html=coverage.out