REPO_ROOT := $(shell git rev-parse --show-toplevel)

VERSION := $(shell git describe --tags --always --dirty)

include $(REPO_ROOT)/makefiles/_python.mk

build-container: ## Build Docker container image
	bazel run --stamp --embed_label=$(VERSION) :load

release-container: ## Upload Docker image to GitHub Container Registry
	bazel run --stamp --embed_label=$(VERSION) :push
