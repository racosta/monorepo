load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "sync",
    srcs = ["sync.go"],
    importpath = "github.com/racosta/monorepo/projects/go_learn_go_with_tests/sync",
    visibility = ["//visibility:public"],
)

go_test(
    name = "sync_test",
    srcs = ["sync_test.go"],
    embed = [":sync"],
)
