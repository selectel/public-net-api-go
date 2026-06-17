default: build

build b:
	go build ./...

fmt f:
	golangci-lint fmt ./...

lint l:
	golangci-lint run ./...

test t:
	go test ./...

cover test_cover:
	go test ./... --cover

.PHONY: lint cover build test fmt
