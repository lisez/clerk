-include .env

PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)


# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get

## build: Compile the binary.
build:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(PROJECTNAME) $(GOFILES)

## go-remod: Write dependencies into go.mod.
go-remod:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod tidy

.PHONY: help install buidl go-remod
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
