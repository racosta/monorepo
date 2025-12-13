TEST_OUTPUT_OPTIONS := --test_output=all

.PHONY: help

# <target_name>: ## <Help text for the target>
# For example:
# build: ## Compiles the project
# clean: ## Removes all generated files

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run first Go binary
	bazel query "kind('go_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel run

test: ## Run bazel test
	bazel test ${TEST_OUTPUT_OPTIONS} ...

benchmark: ## Run bazel benchmark for Go tests
	bazel query "kind('go_test', ...)" --output=label | \
		head -n 1 | \
		xargs -I{} bazel run {} -- -test.bench=. -test.benchmem
