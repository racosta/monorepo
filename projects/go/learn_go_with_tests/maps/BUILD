load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "maps_lib",
    srcs = ["dictionary.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/maps",
    visibility = ["//visibility:public"],
)

go_test(
    name = "maps_test",
    srcs = ["dictionary_test.go"],
    embed = [":maps_lib"],
)
