.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: all test build run
SHELL := /bin/bash
PKG := github.com/flyinprogrammer/pullhashi
PKGS := $(shell go list ./... | grep -v /vendor | grep -v db | grep -v /mock | grep -v /slackapi)
EXECUTABLE := $(shell basename $(PKG))
VERSION := $(shell cat VERSION)
BUILDS := \
	build/$(EXECUTABLE)-v$(VERSION)-darwin-amd64 \
	build/$(EXECUTABLE)-v$(VERSION)-linux-amd64 \
	build/$(EXECUTABLE)-v$(VERSION)-windows-amd64

all: build

build/$(EXECUTABLE)-v$(VERSION)-darwin-amd64:
	GOARCH=amd64 GOOS=darwin go build -o "$@/$(EXECUTABLE)" ./cmd/pullhashi/main.go
build/$(EXECUTABLE)-v$(VERSION)-linux-amd64:
	GOARCH=amd64 GOOS=linux go build -o "$@/$(EXECUTABLE)" ./cmd/pullhashi/main.go
build/$(EXECUTABLE)-v$(VERSION)-windows-amd64:
	GOARCH=amd64 GOOS=windows go build -o "$@/$(EXECUTABLE).exe" ./cmd/pullhashi/main.go

build: $(BUILDS)
