load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_hello_world",
    srcs = ["hello_world.go"],
    importpath = "github.com/racosta/monorepo/projects/go_hello_world",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_hello_world_test",
    srcs = ["hello_world_test.go"],
    embed = [":go_hello_world"],
)
