REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_common.mk

.PHONY: all build run

all: build ## Default target (build)

build: ## Build the first C++ binary
	bazel query "kind('cc_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel build

run: ## Run first C++ binary
	bazel query "kind('cc_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel run
