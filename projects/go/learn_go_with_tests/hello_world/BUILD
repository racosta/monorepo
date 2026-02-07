load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "hello_world_lib",
    srcs = ["hello.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/hello_world",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "hello_world",
    embed = [":hello_world_lib"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "hello_world_test",
    srcs = ["hello_test.go"],
    embed = [":hello_world_lib"],
)
