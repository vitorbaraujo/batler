GO ?= go
GOLANGCI_LINT := $(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run