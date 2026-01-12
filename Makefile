BINARY=devctl

.PHONY: build run test fmt vet clean check

build:
	go build -o $(BINARY) .

run:
	go run .

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet test

clean:
	go clean
