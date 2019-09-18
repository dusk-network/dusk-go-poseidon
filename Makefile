PROJECT_NAME := "dusk-go-poseidon"
PKG := "github.com/dusk-network/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
TEST_LIST := $(shell go list ${PKG}/...)
#TEST_FLAGS := "-count=1"
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
.PHONY: all fmt lintdep lint testdep test coverage coverhtml testclean clean help
all: lint test
fmt: ## Format the go files
	@gofmt -w ${GO_FILES}
lintdep: ## Get the dependencies for the lint
	@go get -u golang.org/x/lint/golint
lint: lintdep ## Lint the files
	@golint -set_exit_status ${PKG_LIST}
testdep: ## Get the dependencies for the tests
	@go get ${PKG_LIST}
test: testdep ## Run unittests
	@go test -p 1 -race -short ${TEST_LIST}
coverage: ## Generate global code coverage report
	chmod u+x ./scripts/coverage.sh
	./scripts/coverage.sh;
coverhtml: ## Generate global code coverage report in HTML
	chmod u+x ./scripts/coverage.sh
	./scripts/coverage.sh html;
testclean: ## Clean the go test cache
	@go clean -testcache
clean: testclean ## Remove previous build
	@rm -rf vendor/*
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
