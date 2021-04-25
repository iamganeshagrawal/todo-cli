.PHONY: build

VERSION ?= v0.0.1
VERSION_PKG_PATH ?= github.com/iamganeshagrawal/todo-cli/pkg/version

build:
	@echo Building the executable binary...
	@go build -o bin/todo.exe main.go
	@echo Build successfully completed

version-build:
	@echo Building $(VERSION) ...
	@go build -v -ldflags="-X '$(VERSION_PKG_PATH).Version=$(VERSION)' -X '$(VERSION_PKG_PATH).BuildTime=25 April 2021'" -o bin/todo.exe
	@echo Build successfully completed