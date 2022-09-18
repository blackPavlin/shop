.PHONY: build
build:
	go build -v ./app/cmd/shop

.PHONY: generate	
generate:
	go generate ./...

.PHONY: openapi/generate
openapi/generate:
	oapi-codegen -generate types -o ./app/internal/controllers/rest/types.gen.go -package rest ./api/openapi.v1.yaml	

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
