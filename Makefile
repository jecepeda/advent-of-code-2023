.PHONY: test build

build:
	CGO_ENABLED=0 go build .

test:
	go test -v ./...