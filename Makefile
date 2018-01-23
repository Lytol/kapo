BINARY_NAME := kapo

all: build

build:
	go build -o $(BINARY_NAME) -v cmd/kapo/*

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)
