.PHONY: help

# <target_name>: ## <Help text for the target>
# For example:
# build: ## Compiles the project
# clean: ## Removes all generated files

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run first Go binary
	bazel query "kind('go_binary', ...)" --output=label | head -n 1 | xargs bazel run
