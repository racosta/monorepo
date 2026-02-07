load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "hello_world",
    srcs = ["hello_world.go"],
    importpath = "github.com/racosta/monorepo/projects/go/hello_world",
    visibility = ["//visibility:public"],
)

go_test(
    name = "hello_world_test",
    srcs = ["hello_world_test.go"],
    embed = [":hello_world"],
)
