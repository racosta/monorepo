load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "concurrency",
    srcs = ["check_websites.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/concurrency",
    visibility = ["//visibility:public"],
)

go_test(
    name = "concurrency_test",
    srcs = [
        "check_websites_benchmark_test.go",
        "check_websites_test.go",
    ],
    embed = [":concurrency"],
)
