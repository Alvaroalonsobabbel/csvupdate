.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint: 
	golangci-lint run

.PHONY: check
check: lint test

.PHONY: build
build: 
	go build -o csvupdate cmd/*.go