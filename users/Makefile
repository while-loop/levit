PROJECT := "levit"
SERVICE := "users"
VERSION := "1"
IMAGE_NAME := ${PROJECT}-${SERVICE}:${VERSION}

.PHONY: proto test get ci docker

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	go generate ./...

test:
	go test -v ./... -cover

get:
	go get -u -t ./...

ci: get test

build:
	docker build -t ${IMAGE_NAME} .

build-nc:
	docker build --no-cache -t ${IMAGE_NAME} .

deploy:
	docker push ${IMAGE_NAME}