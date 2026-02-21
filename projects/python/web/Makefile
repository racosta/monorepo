REPO_ROOT := $(shell git rev-parse --show-toplevel)

VERSION := $(shell git describe --tags --always --dirty)

include $(REPO_ROOT)/makefiles/_python.mk

build-container: ## Build Docker container
	$(BAZEL) run --stamp --embed_label=$(VERSION) :load
