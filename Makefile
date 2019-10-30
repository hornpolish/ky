PROJECT_NAME := ky
PKG := github.com/hornpolish/$(PROJECT_NAME)
PKG_LIST := $(shell go list ${PKG}/...)
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)
VERSION  := $(shell cat VERSION)
DATETIME := $(shell date +%FT%T%z) 
GITHASH  := $(shell git rev-parse HEAD | head -c 8)
 
.PHONY: all dep lint vet test test-coverage build clean
 
all: build

dep: ## Get the dependencies
	go mod download

lint: ## Lint Golang files
	golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	go vet ${PKG_LIST}

test: ## Run unittests
	go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	cat cover.out >> coverage.txt

build: dep ## Build the binary file
	go build -ldflags "-X ${PKG}/cmd.version=${VERSION} -X ${PKG}/cmd.date=${DATETIME} -X ${PKG}/cmd.commit=${GITHASH}" -o build $PKG
 
clean: ## Remove previous build
	rm -f $(PROJECT_NAME)/build
 
help: ## Display this help screen
	grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
