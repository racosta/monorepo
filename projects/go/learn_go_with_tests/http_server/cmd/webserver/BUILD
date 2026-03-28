load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "webserver_lib",
    srcs = ["main.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/cmd/webserver",
    visibility = ["//visibility:private"],
    deps = ["//projects/go/learn_go_with_tests/http_server"],
)

go_binary(
    name = "webserver",
    embed = [":webserver_lib"],
    visibility = ["//visibility:public"],
)
