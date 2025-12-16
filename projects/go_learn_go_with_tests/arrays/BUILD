load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "arrays",
    srcs = ["sum.go"],
    importpath = "github.com/racosta/monorepo/projects/go_learn_go_with_tests/arrays",
    visibility = ["//visibility:public"],
)

go_test(
    name = "arrays_test",
    srcs = ["sum_test.go"],
    embed = [":arrays"],
)
