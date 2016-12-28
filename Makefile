.PHONY: build deps update test

NAME := genetemp
VERSION := $(shell git describe --tags --abbrev=0)

## setup
setup:
	go get github.com/Masterminds/glide
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golang/lint/golint

deps:
	glide install
update:
	glide update

test: deps
	go test $$(glide novendor)

build:
	go build -ldflags "-X 'main.version=$(VERSION)'"
