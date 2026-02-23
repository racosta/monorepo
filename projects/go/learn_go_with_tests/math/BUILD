load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "math",
    srcs = ["clockface.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/math",
    visibility = ["//visibility:public"],
)

go_test(
    name = "math_test",
    srcs = [
        "clockface_acceptance_test.go",
        "clockface_test.go",
    ],
    embed = [":math"],
)

go_test(
    name = "math_acceptance_test",
    srcs = ["clockface_acceptance_test.go"],
    embed = [":math"],
)
