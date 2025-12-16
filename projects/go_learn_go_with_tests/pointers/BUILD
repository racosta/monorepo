load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "pointers",
    srcs = ["wallet.go"],
    importpath = "github.com/racosta/monorepo/projects/go_learn_go_with_tests/pointers",
    visibility = ["//visibility:public"],
)

go_test(
    name = "pointers_test",
    srcs = ["wallet_test.go"],
    embed = [":pointers"],
)
