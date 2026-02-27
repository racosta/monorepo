REPO_ROOT := $(shell git rev-parse --show-toplevel)

include $(REPO_ROOT)/makefiles/_go.mk

.PHONY: clockface

clockface: ## Run the clockface Go binary
	bazel run clockface:clockface > clock.svg
