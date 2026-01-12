BINARY=devctl

.PHONY: build run mock test fmt vet clean check

build:
	go build -o $(BINARY) .

run:
	go run .

mock:
	go run ./cmd/mockserver

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet test

clean:
	go clean
