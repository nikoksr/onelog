export GO111MODULE := on
export GOPROXY ?= https://proxy.golang.org,direct

###############################################################################
# DEPENDENCIES
###############################################################################

install:
	@go install mvdan.cc/gofumpt@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
.PHONY: install

###############################################################################
# TESTS
###############################################################################

test:
	go test -failfast -race ./...
.PHONY: test

gen-coverage:
	@go test -race -covermode=atomic -coverprofile=coverage.out ./... > /dev/null
.PHONY: gen-coverage

coverage: gen-coverage
	go tool cover -func coverage.out
.PHONY: coverage

coverage-html: gen-coverage
	go tool cover -html=coverage.out -o cover.html
.PHONY: coverage-html

###############################################################################
# CODE HEALTH
###############################################################################

fmt:
	@gofumpt -w -l .
	@goimports -local github.com/nikoksr/onelog -w -l .
.PHONY: fmt

lint:
	@golangci-lint run --config .golangci.yml
.PHONY: lint

ci: lint test
.PHONY: ci

###############################################################################

.DEFAULT_GOAL := ci
