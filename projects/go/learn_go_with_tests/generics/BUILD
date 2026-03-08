load("@rules_go//go:def.bzl", "go_library", "go_test")

go_test(
    name = "generics_test",
    srcs = ["generics_test.go"],
    embed = [":generics"],
)

go_library(
    name = "generics",
    srcs = [
        "assert.go",
        "stack.go",
    ],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/generics",
    visibility = ["//visibility:public"],
)
