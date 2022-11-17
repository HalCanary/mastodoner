all: follow listfollows poststatus

.PHONY: all, clean, gofmt

follow listfollows poststatus: %: $(shell find . -name '*.go')
	go build ./cmd/$*

clean:
	git clean -fx

gofmt:
	gofmt -w */*.go cmd/*/*.go
