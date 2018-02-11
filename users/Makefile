# service specific vars
SERVICE     := users
VERSION     := 0.0.1

ORG         := levitgo
TARGET      := ${SERVICE}d
COMMIT      := $(shell git rev-parse --short HEAD)
BUILD_TIME  := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
IMAGE_NAME  := ${ORG}/${SERVICE}
PACKAGE 	:= $(shell pwd | sed "s,${GOPATH}/src/,,")

.PHONY: proto deps test build cont cont-nc all deploy help clean lint
.DEFAULT_GOAL := help

help: ## halp
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: deps lint test build ## get && test && build

build: clean lint ## build service binary file
	@echo "[build] building go binary"
	@go build \
		-ldflags "-s -w \
		-X ${PACKAGE}/version.Version=${VERSION} \
		-X ${PACKAGE}/version.Commit=${COMMIT} \
		-X ${PACKAGE}/version.BuildTime=${BUILD_TIME}" \
		-o ${GOPATH}/bin/${TARGET} ./cmd/${TARGET}
	${SERVICE}d -v

clean: ## remove service bin from $GOPATH/bin
	@echo "[clean] removing service files"
	rm -f ${GOPATH}/bin/${TARGET}

cont: ## build a non-cached service container
	docker build -t ${IMAGE_NAME} -t ${IMAGE_NAME}:${VERSION} . --no-cache

cont-c: ## build a cached service container
	docker build -t ${IMAGE_NAME} -t ${IMAGE_NAME}:${VERSION} .

deploy: ## deploy lastest built container to docker hub
	docker push ${IMAGE_NAME}

deps: ## get service pkg + test deps
	@echo "[deps] getting go deps"
	go get -v -t ./...

lint: ## apply golint
	@echo "[lint] applying go fmt & vet"
	go fmt ./...
	go vet ./...

proto: ## generate protobuf go objects
	go get github.com/golang/protobuf/protoc-gen-go
	go generate ./...

release: test cont deploy ## build and deploy a docker container

test: lint ## test service code
	@echo "[test] running tests w/ cover"
	go test ./... -cover