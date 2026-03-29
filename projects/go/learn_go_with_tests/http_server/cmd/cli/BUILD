load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "cli_lib",
    srcs = ["main.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/cmd/cli",
    visibility = ["//visibility:private"],
    deps = ["//projects/go/learn_go_with_tests/http_server"],
)

go_binary(
    name = "cli",
    data = ["game.db.json"],
    embed = [":cli_lib"],
    visibility = ["//visibility:public"],
)
