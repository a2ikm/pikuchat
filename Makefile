.PHONY: all
all: build install-tools

.PHONY: build
build:
	go build .

.PHONY: install-tools
install-tools:
	go get github.com/cosmtrek/air
