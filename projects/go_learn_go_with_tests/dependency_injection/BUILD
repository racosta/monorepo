load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "dependency_injection_lib",
    srcs = ["di.go"],
    importpath = "github.com/racosta/monorepo/projects/go_learn_go_with_tests/dependency_injection",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "dependency_injection_bin",
    embed = [":dependency_injection_lib"],
)

go_test(
    name = "dependency_injection_test",
    srcs = ["di_test.go"],
    embed = [":dependency_injection_lib"],
)
