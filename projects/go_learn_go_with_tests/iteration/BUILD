load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "iteration",
    srcs = ["repeat.go"],
    importpath = "github.com/racosta/monorepo/projects/go_learn_go_with_tests/iteration",
    visibility = ["//visibility:public"],
)

go_test(
    name = "iteration_test",
    srcs = ["repeat_test.go"],
    embed = [":iteration"],
)
