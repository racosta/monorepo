# Use bazelisk if available, otherwise fallback to bazel
BAZEL := $(shell command -v bazelisk 2> /dev/null || echo bazel)

TEST_OUTPUT_OPTIONS := --test_output=all

.PHONY: help gomod-init gowork-use gazelle all build run test coverage coverage-report benchmark clean pristine

REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_common.mk

gomod-init: ## Initialize a Go module using current directory
	$(BAZEL) run @rules_go//go mod init github.com/racosta/monorepo/$(shell git rev-parse --show-prefix | sed 's,/$$,,')

gomod-tidy: ## Tidy the Go module
	$(BAZEL) run @rules_go//go mod tidy

gowork-use: ## Add current Go module to the Go workspace
	cd $(shell git rev-parse --show-toplevel) && \
		$(BAZEL) run @rules_go//go work use $(shell git rev-parse --show-prefix | sed 's,/$$,,')

gazelle: ## Run gazelle on current directory
	cd $(shell git rev-parse --show-toplevel) && \
		$(BAZEL) run //:gazelle -- $(shell git rev-parse --show-prefix | sed 's,/$$,,')

all: build ## Default target (build)

build: ## Build the first Go binary
	$(BAZEL) query "kind('go_binary', ...)" --output=label | \
		head -n 1 | \
		xargs $(BAZEL) build

run: ## Run first Go binary
	$(BAZEL) query "kind('go_binary', ...)" --output=label | \
		head -n 1 | \
		xargs $(BAZEL) run

test: ## Run bazel test
	$(BAZEL) test ${TEST_OUTPUT_OPTIONS} ...

coverage: ## Run bazel coverage
	$(BAZEL) coverage ${TEST_OUTPUT_OPTIONS} --combined_report=lcov ...

coverage-report: coverage ## Generate coverage report HTML
	cd $(shell git rev-parse --show-toplevel) && \
		genhtml --branch-coverage --legend --num-spaces=2 --output genhtml "$(shell $(BAZEL) info output_path)/_coverage/_coverage_report.dat"

benchmark: ## Run bazel benchmark for Go tests
	$(BAZEL) query "kind('go_test', ...)" --output=label | \
		head -n 1 | \
		xargs -I{} $(BAZEL) run {} -- -test.bench=. -test.benchmem
