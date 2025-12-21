load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "select",
    srcs = ["racer.go"],
    importpath = "github.com/racosta/monorepo/projects/go_learn_go_with_tests/select",
    visibility = ["//visibility:public"],
)

go_test(
    name = "select_test",
    srcs = ["racer_test.go"],
    embed = [":select"],
)
