GO ?= go
GOLANGCI_LINT := $(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: check
check:
	$(GO) test ./... -coverpkg=./... -coverprofile=coverage.out

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run

.PHONY: cover
cover:
	$(GO) tool cover -func=coverage.out