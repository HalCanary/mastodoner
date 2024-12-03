## Go Makefile
## Copyright 2022 Hal Canary.  See LICENSE.md

go_commands = $(notdir $(wildcard cmd/*))
go_binaries = $(addprefix build/,${go_commands})

all: build test

build: ${go_binaries}

build/dependencies.stamp: go.mod
	go get ./...
	@mkdir -p build
	@touch $@

clean:
	rm -rf build

fmt:
	gofmt -w */*.go cmd/*/*.go

test: build/dependencies.stamp
	go test ./...

install: ${go_binaries}
	mkdir -p ~/bin
	mv $^ ~/bin

define GoCommandTemplate
build/$1: $$(wildcard cmd/$1/*.go wildcard */*.go) build/dependencies.stamp
	go build -o build ./cmd/$1
endef
$(foreach d,$(go_commands),$(eval $(call GoCommandTemplate,$d)))

.PHONY: all, build, clean, fmt, test, install
