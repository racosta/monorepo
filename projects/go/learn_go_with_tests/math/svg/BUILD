load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "svg",
    srcs = ["svg.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/math/svg",
    visibility = ["//visibility:public"],
    deps = ["//projects/go/learn_go_with_tests/math"],
)

go_test(
    name = "svg_test",
    srcs = ["svg_test.go"],
    embed = [":svg"],
)
