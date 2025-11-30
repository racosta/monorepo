load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_quiz_lib",
    srcs = ["main.go"],
    importpath = "github.com/racosta/monorepo/" + package_name(),
    visibility = ["//visibility:private"],
)

go_binary(
    name = "go_quiz",
    data = ["problems.csv"],
    embed = [":go_quiz_lib"],
    visibility = ["//visibility:public"],
)
