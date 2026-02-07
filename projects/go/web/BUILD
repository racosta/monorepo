load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "web",
    embed = [":web_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "web_lib",
    srcs = ["main.go"],
    importpath = "github.com/racosta/monorepo/projects/go/web",
    visibility = ["//visibility:private"],
    deps = [
        "//projects/go/hello_world",
        "@com_github_gin_gonic_gin//:gin",
    ],
)
