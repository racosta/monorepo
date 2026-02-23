load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "clockface_lib",
    srcs = ["main.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/math/clockface",
    visibility = ["//visibility:private"],
    deps = ["//projects/go/learn_go_with_tests/math"],
)

go_binary(
    name = "clockface",
    embed = [":clockface_lib"],
    visibility = ["//visibility:public"],
)
