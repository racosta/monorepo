# Use bazelisk if available, otherwise fallback to bazel
BAZEL := $(shell command -v bazelisk 2> /dev/null || echo bazel)
TEST_OUTPUT_OPTIONS := --test_output=errors

.PHONY: help all build run test clean pristine

REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_common.mk

all: build ## Default target (build)

build: ## Build the first Python binary
	$(BAZEL) query "kind('py_binary', ...)" --output=label | \
		head -n 1 | \
		xargs $(BAZEL) build

run: ## Run first Python binary
	$(BAZEL) query "kind('py_binary', ...)" --output=label | \
		head -n 1 | \
		xargs $(BAZEL) run

test: ## Run bazel test
	$(BAZEL) test ... ${TEST_OUTPUT_OPTIONS}
