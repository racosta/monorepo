load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "arrays",
    srcs = [
        "bad_bank.go",
        "collection_fun.go",
        "sum.go",
    ],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/arrays",
    visibility = ["//visibility:public"],
)

go_test(
    name = "arrays_test",
    srcs = [
        "assert_test.go",
        "bad_bank_test.go",
        "sum_test.go",
    ],
    embed = [":arrays"],
)
