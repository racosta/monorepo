TEST_OUTPUT_OPTIONS := --test_output=all

.PHONY: help gomod-init gowork-use gazelle all build run test coverage coverage-report benchmark clean pristine

# <target_name>: ## <Help text for the target>
# For example:
# build: ## Compiles the project
# clean: ## Removes all generated files

help: ## Show the help with the list of commands
	@sort $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[0;33m%s\033[0m\n", substr($$0, 5) } '
	@echo ""

gomod-init: ## Initialize a Go module using current directory
	bazel run @rules_go//go mod init github.com/racosta/monorepo/$(shell git rev-parse --show-prefix | sed 's,/$$,,')

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

test: ## Run bazel test
	bazel test ${TEST_OUTPUT_OPTIONS} ...

coverage: ## Run bazel coverage
	bazel coverage ${TEST_OUTPUT_OPTIONS} --combined_report=lcov ...

coverage-report: coverage ## Generate coverage report HTML
	cd $(shell git rev-parse --show-toplevel) && \
		genhtml --branch-coverage --legend --num-spaces=2 --output genhtml "$(shell bazel info output_path)/_coverage/_coverage_report.dat"

benchmark: ## Run bazel benchmark for Go tests
	bazel query "kind('go_test', ...)" --output=label | \
		head -n 1 | \
		xargs -I{} bazel run {} -- -test.bench=. -test.benchmem

clean: ## Clean Bazel output files
	bazel clean

pristine: ## Clean and purge all Bazel files
	bazel clean --expunge
