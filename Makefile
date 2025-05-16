.DEFAULT_GOAL := help

vendor      := vendor
target      := target
bin         := $(target)/bin
reports     := $(target)/reports

## build: Compile binaries.
go_src := $(shell find * -name *.go -not -path "$(vendor)/*" -not -path "$(target)/*")
go_out := $(patsubst cmd/%/main.go,$(bin)/%,$(wildcard cmd/*/main.go))

.PHONY: build
build: $(go_out)

$(bin)/%: cmd/%/main.go $(go_src) | $(bin)
	@go build --trimpath -o=$@ $<

$(bin):
	@mkdir -p $@

$(reports):
	@mkdir -p $@

## generate: Run generators.
.PHONY: generate
generate: go/generate openapi/generate

.PHONY: go/generate
go/generate:
	@go generate ./...

.PHONY: openapi/generate
openapi/generate:
	@go tool oapi-codegen -config ./api/codegen.config.yaml ./api/openapi.v1.yaml

## tests: Run tests.
.PHONY: test tests
test tests: go/test

.PHONY: go/test
go/test: $(go_src) | $(reports)
	@go test -v -covermode=atomic -coverprofile=$(reports)/cover.out ./...

## lint: Run static analysis.
.PHONY: lint
lint: go/lint

.PHONY: go/lint
go/lint:
	@golangci-lint run

## help: Display available targets.
.PHONY: help
help: Makefile
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@sed -n 's/^## //p' $< | awk -F ':' '{printf "  %-20s%s\n",$$1,$$2}'
