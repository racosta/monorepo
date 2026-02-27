REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_common.mk

.PHONY: all build run

all: build ## Default target (build)

build: ## Build the first Rust binary
	bazel query "kind('rust_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel build

run: ## Run first Rust binary
	bazel query "kind('rust_binary', ...)" --output=label | \
		head -n 1 | \
		xargs bazel run

gen_rust_project: ## Generate rust-project.json for Rust analyzer
	bazel run @rules_rust//tools/rust_analyzer:gen_rust_project -- \
		--workspace=$(REPO_ROOT)

gen_rust_launch_json: ## Generate VSCode launch.json for Rust debugging
	bazel run @rules_rust//tools/vscode:gen_launch_json -- \
		--workspace-root=$(REPO_ROOT) \
		--output=$(REPO_ROOT)/.vscode/launch.json
