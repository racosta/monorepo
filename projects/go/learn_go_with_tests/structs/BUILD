load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "structs",
    srcs = ["shapes.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/structs",
    visibility = ["//visibility:public"],
)

go_test(
    name = "structs_test",
    srcs = ["shapes_test.go"],
    embed = [":structs"],
)
