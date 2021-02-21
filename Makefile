-include .env

PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin


# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install:
	@GOBIN=$(GOBIN) go get

## build: Compile the binary.
build:
	@GOBIN=$(GOBIN) go build -ldflags="-s -w" -o $(PROJECTNAME)

## go-remod: Write dependencies into go.mod.
go-remod:
	@GOBIN=$(GOBIN) go mod tidy

## go-dev: Test run on development
go-dev:
	@GOBIN=$(GOBIN) go run main.go

## go-test: Runing testings
go-test:
	@GOBIN=$(GOBIN) go test -v ./...

.PHONY: help install buidl go-remod go-dev
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
