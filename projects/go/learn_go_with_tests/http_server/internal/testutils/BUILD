load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "testutils",
    srcs = ["helpers.go"],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils",
    visibility = ["//projects/go/learn_go_with_tests/http_server:__subpackages__"],
    deps = [
        "//projects/go/learn_go_with_tests/http_server/internal/league",
        "//projects/go/learn_go_with_tests/http_server/internal/player:player_lib",
    ],
)
