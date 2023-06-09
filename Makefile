GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
## fully-qualified path to this Makefile
MKFILE_PATH := $(realpath $(lastword $(MAKEFILE_LIST)))

## fully-qualified path to the current directory
CURRENT_DIR := $(patsubst %/,%,$(dir $(MKFILE_PATH)))
CUR_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BINARY = "${CUR_DIR}/.bin/apis"
DOCKER_TAG=apis:latest

default: help

.PHONY: help
## This help screen
help:
	@printf "Available targets:\n\n"
	@awk '/^[a-zA-Z\-\_0-9%:\\]+/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
          if (helpMessage) { \
            helpCommand = $$1; \
            helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
      gsub("\\\\", "", helpCommand); \
      gsub(":+$$", "", helpCommand); \
            printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
          } \
        } \
        { lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"

all: build test

run:
	@go run main.go server
.PHONY: run

up:
	@docker-compose up
.PHONY: up

stop:
	@docker-compose down -
.PHONY: stop

clean:
	@go clean

gen-api:
	@buf generate && cp third_party/OpenAPI/mengoaingu.swagger.json swagger_ui/mengoaingu.swagger.json

swagger-ui:
	@python3 -m http.server --directory ./swagger_ui 9900

.PHONY: gen-api swagger-ui

## Build the binary for linux
linux:
	@GOOS=linux GOARCH=${GOARCH} CGO_ENABLED=0 go build ${LDFLAGS} -o ${BINARY} .

## Sync the binary to the remote server
sync:
	scp -i  ~/.ssh/id_rsa.pub ${BINARY} mengoaingu:/root/app/apis

docker-stack:
	docker stack deploy -c docker-stack.yaml mengoaingu
	@GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY} .

