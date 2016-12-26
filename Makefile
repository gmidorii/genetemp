.PHONY: build

NAME := genetemp
VERSION := $(shell git describe --tags --abbrev=0)

deps:
	glide install
update:
	glide update
build:
	go build -ldflags "-X 'main.version=$(VERSION)'"
