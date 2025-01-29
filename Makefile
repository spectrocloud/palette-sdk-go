# If you update this file, please follow:
# https://suva.sh/posts/well-documented-makefiles/

.DEFAULT_GOAL:=help

# Go vars
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Output
TIME   = `date +%H:%M:%S`
GREEN  := $(shell printf "\033[32m")
RED    := $(shell printf "\033[31m")
CNone  := $(shell printf "\033[0m")
OK   = echo ${TIME} ${GREEN}[ OK ]${CNone}
ERR  = echo ${TIME} ${RED}[ ERR ]${CNone} "error:"

##@ Help Targets

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[0m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build Targets

generate: ## Generate models
	(cd api && ./generate.sh ./)

##@ Static Analysis Targets

check-diff: reviewable ## Execute branch is clean
	git --no-pager diff
	git diff --quiet || ($(ERR) please run 'make reviewable' to include all changes && false)
	@$(OK) branch is clean

reviewable: fmt vet lint ## Ensure code is ready for review
	git submodule update --remote
	go mod tidy

fmt: ## Run go fmt against code
	go fmt ./...

vet: ## Run go vet against code
	go vet ./...

lint: golangci-lint ## Run golangci-lint against code
	$(GOLANGCI_LINT) run -v ./...

#pre-commit-install: pre-commit ## Install pre-commit hooks
#	@if [ "$(GITHUB_ACTIONS)" != "true" ]; then \
#		pre-commit install --hook-type commit-msg; \
#		pre-commit install --hook-type pre-commit; \
#	fi

##@ Test Targets

.PHONY: test
test: ## Run acceptance tests
	go test -covermode=atomic -coverpkg=./... -coverprofile=profile.cov ./... -timeout 10m

##@ Tools Targets

BIN_DIR ?= ./bin
bin-dir:
	test -d $(BIN_DIR) || mkdir $(BIN_DIR)

GOLANGCI_VERSION ?= 1.62.2
.PHONY: golangci-lint
golangci-lint: bin-dir
	if ! test -f $(BIN_DIR)/golangci-lint-linux-amd64; then \
		curl -LOs https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_VERSION)/golangci-lint-$(GOLANGCI_VERSION)-linux-amd64.tar.gz; \
		tar -zxf golangci-lint-$(GOLANGCI_VERSION)-linux-amd64.tar.gz; \
		mv golangci-lint-$(GOLANGCI_VERSION)-*/golangci-lint $(BIN_DIR)/golangci-lint-linux-amd64; \
		chmod +x $(BIN_DIR)/golangci-lint-linux-amd64; \
		rm -rf ./golangci-lint-$(GOLANGCI_VERSION)-linux-amd64*; \
	fi
	if ! test -f $(BIN_DIR)/golangci-lint-$(GOOS)-$(GOARCH); then \
		curl -LOs https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_VERSION)/golangci-lint-$(GOLANGCI_VERSION)-$(GOOS)-$(GOARCH).tar.gz; \
		tar -zxf golangci-lint-$(GOLANGCI_VERSION)-$(GOOS)-$(GOARCH).tar.gz; \
		mv golangci-lint-$(GOLANGCI_VERSION)-*/golangci-lint $(BIN_DIR)/golangci-lint-$(GOOS)-$(GOARCH); \
		chmod +x $(BIN_DIR)/golangci-lint-$(GOOS)-$(GOARCH); \
		rm -rf ./golangci-lint-$(GOLANGCI_VERSION)-$(GOOS)-$(GOARCH)*; \
	fi
GOLANGCI_LINT=$(BIN_DIR)/golangci-lint-$(GOOS)-$(GOARCH)

.PHONY: pre-commit
pre-commit:
	@if [ "$(GITHUB_ACTIONS)" != "true" ]; then \
		@command -v pre-commit >/dev/null 2>&1 || { \
			echo "pre-commit not found, downloading..."; \
			pip install pre-commit; \
		} \
	fi
