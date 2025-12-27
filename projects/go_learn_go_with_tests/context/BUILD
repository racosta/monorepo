load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "context",
    srcs = ["context.go"],
    importpath = "github.com/racosta/monorepo/projects/go_learn_go_with_tests/context",
    visibility = ["//visibility:public"],
)

go_test(
    name = "context_test",
    srcs = ["context_test.go"],
    embed = [":context"],
)
