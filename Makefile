MAIN_PACKAGE_PATH := ./cmd/test
BINARY_NAME := test

.PHONY: default
default: build

.PHONY: tidy
tidy:
	go fmt ./...

.PHONY: build
build: tidy
	go build -o=$(CURDIR)/dist/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: clean
clean:
	@rm -rf $(CURDIR)/dist
