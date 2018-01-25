BINARY_NAME := kapo
DEFAULT_DATABASE := kapo.db

all: build

build:
	go build -o $(BINARY_NAME) -v cmd/kapo/*

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(DEFAULT_DATABASE)
