
default: test

.PHONY: test
test:
	go test -race ./...

.PHONY: install
install:
	go install .
	bearings install -s zsh
	bearings install -s bash
	bearings install -s fish


