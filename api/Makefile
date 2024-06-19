GORUN := go run
GOBIN := $(shell go env GOPATH)/bin
GOLINT := $(GOBIN)/golint
GOGET := go get
DC := docker compose

.PHONY: deps
deps:
	$(GOGET) -u ./...

.PHONY: install-tools
install-tools:
	@echo "Installing tools..."
	@GOBIN=$(GOBIN) go install golang.org/x/lint/golint@latest

.PHONY: run
run:
	$(DC) up -d --remove-orphans

.PHONY: down
down:
	$(DC) rm -svf

.PHONY: lint
lint: install-tools
	$(GOLINT) ./...
