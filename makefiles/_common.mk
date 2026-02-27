TEST_OUTPUT_OPTIONS := --test_output=errors

.PHONY: help clean pristine test

# <target_name>: ## <Help text for the target>
# For example:
# build: ## Compiles the project
# clean: ## Removes all generated files

help: ## Show the help with the list of commands
	@sort $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[0;33m%s\033[0m\n", substr($$0, 5) } '
	@echo ""

clean: ## Clean Bazel output files
	bazel clean

pristine: ## Clean and purge all Bazel files
	bazel clean --expunge

test: ## Run bazel test
	bazel test ${TEST_OUTPUT_OPTIONS} ...
