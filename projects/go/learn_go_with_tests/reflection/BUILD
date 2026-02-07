load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "reflection",
    srcs = ["reflection.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/reflection",
    visibility = ["//visibility:public"],
)

go_test(
    name = "reflection_test",
    srcs = ["reflection_test.go"],
    embed = [":reflection"],
)
