app := stine

.PHONY: all
all: test build

.PHONY: test
test:
	go vet ./...
	go test -v -vet=off ./...

.PHONY: build
build: export GOVERSION := $(shell go version | awk '{print $$3 " on " $$4;}')
build:
	goreleaser release --rm-dist --snapshot
	cp ./dist/$(app)_linux_amd64/$(app) ./$(app)
	./$(app) -v

.PHONY: clean
clean:
	@rm -rf dist
	@rm $(app)

.PHONY: update
update:
	go get -u
	go mod tidy
