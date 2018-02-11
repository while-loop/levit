.PHONY: proto deps test build cont cont-nc all deploy help clean lint
.DEFAULT_GOAL := help


SERVICES := users hub

help: ## halp
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: ## make all services
	@for dir in $(SERVICES); do \
		make -C $$dir all ; \
	done

clean: ## clean all services
	@for dir in $(SERVICES); do \
		make -C $$dir clean ; \
	done

build: clean lint ## build service binary file
	@for dir in $(SERVICES); do \
		make -C $$dir build ; \
	done

deps: ## get service pkg + test deps
	@for dir in $(SERVICES); do \
		make -C $$dir deps ; \
	done

lint: ## apply golint to all projects
	@for dir in $(SERVICES); do \
		make -C $$dir lint ; \
	done

test: ## test service code
	@for dir in $(SERVICES); do \
		make -C $$dir test ; \
	done
