.PHONY: all build clean

all: build-server build-lambda

build-server:
	go build -o bin/server ./cmd/server

build-lambda:
	GOOS=linux GOARCH=amd64 go build -o bin/bootstrap ./cmd/lambda

zip-lambda: build-lambda
	cd bin && zip -j lambda.zip bootstrap

build-and-zip-lambda:
	GOOS=linux GOARCH=amd64 go build -o bin/bootstrap ./cmd/lambda
	cd bin && zip -j lambda.zip bootstrap

clean:
	rm -rf bin/*
