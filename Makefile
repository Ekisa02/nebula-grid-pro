BINARY=nebula-grid

.PHONY: build run test

build:
	go build -o bin/$(BINARY) ./cmd/cloudmesh

run: build
	./bin/$(BINARY)

test:
	go test ./... -v

fmt:
	gofmt -s -w .
