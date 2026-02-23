load("@rules_go//go:def.bzl", "go_library", "go_test")

# gazelle:go_test file

go_library(
    name = "math",
    srcs = ["clockface.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/math",
    visibility = ["//visibility:public"],
)

go_test(
    name = "clockface_test",
    srcs = ["clockface_test.go"],
    embed = [":math"],
)
