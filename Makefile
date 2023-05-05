GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
## fully-qualified path to this Makefile
MKFILE_PATH := $(realpath $(lastword $(MAKEFILE_LIST)))

## fully-qualified path to the current directory
CURRENT_DIR := $(patsubst %/,%,$(dir $(MKFILE_PATH)))
DOCKER_TAG=apis-definitions:latest

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
	@go run main.go server2 
.PHONY: run

up:
	@docker-compose up
.PHONY: up

stop:
	@docker-compose down -
.PHONY: stop

clean:
	@go clean

migrate_up: migrate_profile_up migrate_task_up

migrate_profile_up:
	migrate -path internal/profile/infrastructure/db/migrations -database "mysql://root:123456@tcp(127.0.0.1:3306)/profile?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true" -verbose up

migrate_task_up:
	migrate -path internal/tasks/infrastructure/db/migrations -database "mysql://root:123456@tcp(127.0.0.1:3306)/tasks?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true" -verbose up

.PHONY: migrate_up migrate_profile_up migrate_task_up

wire:
	@cd pkg/bindings/ && wire ./...
.PHONY: wire

gen-api:
	@buf generate
.PHONY: gen-api