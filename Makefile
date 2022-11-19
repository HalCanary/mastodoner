## Go Makefile
## Copyright 2022 Hal Canary.  See LICENSE.md

CMDS = $(notdir $(wildcard cmd/*))

all: $(CMDS)

clean:
	rm -f $(CMDS)

gofmt:
	gofmt -w */*.go cmd/*/*.go

gotest:
	go test ./...

install: $(CMDS)
	mkdir -p ~/bin
	mv $(CMDS) ~/bin

define GoCommandTemplate
$1: $$(wildcard cmd/$1/*.go wildcard */*.go)
	go build ./cmd/$1
endef
$(foreach d,$(CMDS),$(eval $(call GoCommandTemplate,$d)))

.PHONY: all, clean, gofmt, gotest, install
