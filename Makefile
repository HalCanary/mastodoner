poststatus: $(shell find . -name '*.go')
	go build ./cmd/poststatus
