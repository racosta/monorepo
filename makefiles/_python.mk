TEST_OUTPUT_OPTIONS := --test_output=all

.PHONY: help all build run test clean pristine

REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_common.mk

all: build ## Default target (build)

build: ## Build the first Python binary
	bazel query "kind('py_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel build

run: ## Run first Python binary
	bazel query "kind('py_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel run

test: ## Run bazel test
	bazel test ${TEST_OUTPUT_OPTIONS} ...
