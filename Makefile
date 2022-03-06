
default: test

.PHONY: test
test:
	go test -race ./...

.PHONY: install
install:
	go install .
	bearings install

