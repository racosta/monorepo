load("@rules_go//go:def.bzl", "go_library", "go_test")

go_test(
    name = "integers_test",
    srcs = ["adder_test.go"],
    embed = [":integers"],
)

go_library(
    name = "integers",
    srcs = ["adder.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/integers",
    visibility = ["//visibility:public"],
)
