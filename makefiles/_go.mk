REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_common.mk

.PHONY: gomod-init gowork-use gazelle all build run coverage coverage-report benchmark

gomod-init: ## Initialize a Go module using current directory
	bazel run @rules_go//go mod init github.com/racosta/monorepo/$(shell git rev-parse --show-prefix | sed 's,/$$,,')

gomod-tidy: ## Tidy the Go module
	bazel run @rules_go//go mod tidy

gowork-use: ## Add current Go module to the Go workspace
	cd $(shell git rev-parse --show-toplevel) && \
		bazel run @rules_go//go work use $(shell git rev-parse --show-prefix | sed 's,/$$,,')

gazelle: ## Run gazelle on current directory
	cd $(shell git rev-parse --show-toplevel) && \
		bazel run //:gazelle -- $(shell git rev-parse --show-prefix | sed 's,/$$,,')

all: build ## Default target (build)

build: ## Build the first Go binary
	bazel query "kind('go_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel build

run: ## Run first Go binary
	bazel query "kind('go_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel run

coverage: ## Run bazel coverage
	bazel coverage ${TEST_OUTPUT_OPTIONS} --combined_report=lcov ...

coverage-report: coverage ## Generate coverage report HTML
	cd $(shell git rev-parse --show-toplevel) && \
		genhtml --branch-coverage --legend --num-spaces=2 --output genhtml "$(shell bazel info output_path)/_coverage/_coverage_report.dat"

benchmark: ## Run bazel benchmark for Go tests
	bazel query "kind('go_test', ...)" --output=label | \
		head -n 1 | \
		xargs -I{} bazel run {} -- -test.bench=. -test.benchmem
