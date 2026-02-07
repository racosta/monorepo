load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "mocking_lib",
    srcs = ["main.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/mocking",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "mocking_bin",
    embed = [":mocking_lib"],
)

go_test(
    name = "mocking_test",
    srcs = ["countdown_test.go"],
    embed = [":mocking_lib"],
)
