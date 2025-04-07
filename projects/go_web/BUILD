load("@rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "go_web",
    embed = [":go_web_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_web_lib",
    srcs = ["main.go"],
    importpath = "github.com/racosta/monorepo/projects/go_web",
    visibility = ["//visibility:private"],
    deps = [
        "//projects/go_hello_world",
        "@com_github_gin_gonic_gin//:gin",
    ],
)
