.PHONY: build
build:
	go build -v ./app/cmd/shop

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
