BINARY_NAME=ws-app
CURRENT_DIR=$(shell pwd)

export GO111MODULE=on

.PHONY: all build clean lint critic test

all: build

build:
	go build -v -o ${BINARY_NAME} ./cmd

clean:
	rm -f ${BINARY_NAME}

lint:
	golangci-lint run -v

test:
	go test -v ./...

init:
	go mod init

tidy:
	go mod tidy

