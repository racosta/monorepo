load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "roman_numerals",
    srcs = ["numeral.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/roman_numerals",
    visibility = ["//visibility:public"],
)

go_test(
    name = "roman_numerals_test",
    srcs = ["numeral_test.go"],
    embed = [":roman_numerals"],
)
