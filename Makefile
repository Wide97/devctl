BINARY=devctl
DEVCTL_BASE_URL ?= http://localhost:8080
FILE_PATH ?= .
DIR ?= .
CALC_OP ?= add
x ?=
y ?=
X ?= $(x)
Y ?= $(y)
X ?= 1
Y ?= 1

.PHONY: build run mock demo help sys ping file-exists file-ls calc test fmt vet clean check

build:
	go build -o $(BINARY) .

run:
	go run .

mock:
	go run ./cmd/mockserver

demo:
	go run ./cmd/demo

help:
	go run . help

sys:
	DEVCTL_BASE_URL=$(DEVCTL_BASE_URL) go run . sys info

ping:
	DEVCTL_BASE_URL=$(DEVCTL_BASE_URL) go run . ping

file-exists:
	DEVCTL_BASE_URL=$(DEVCTL_BASE_URL) go run . file exists $(FILE_PATH)

file-ls:
	DEVCTL_BASE_URL=$(DEVCTL_BASE_URL) go run . file ls $(DIR)

calc:
	go run . calc $(CALC_OP) $(X) $(Y)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet test

clean:
	go clean
