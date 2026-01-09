BINARY=devctl

.PHONY: build run test fmt vet clean

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

clean:
	go clean
