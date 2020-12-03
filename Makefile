GO ?= go
GOLANGCI_LINT := $(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint
GOLINT := $(GO) run golang.org/x/lint/golint

PACKAGES := $(shell $(GO) list ./...)

.PHONY: check
check:
	$(GO) test ./... -coverpkg=./... -coverprofile=coverage.out

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run
	$(GOLINT) -set_exit_status $(PACKAGES)

.PHONY: cover
cover:
	$(GO) tool cover -func=coverage.out