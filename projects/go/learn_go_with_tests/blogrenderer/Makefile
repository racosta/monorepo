REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_go.mk

approval-tests-update: ## Update approval test golden files
	bazel run :blogrenderer_test -- -update
