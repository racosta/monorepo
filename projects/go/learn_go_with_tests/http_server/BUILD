load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

# gazelle:go_test file

go_library(
    name = "http_server_lib",
    srcs = [
        "file_system_store.go",
        "in_memory_player_store.go",
        "main.go",
        "server.go",
        "tape.go",
    ],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server",
    visibility = ["//visibility:public"],
    deps = [
        "//projects/go/learn_go_with_tests/http_server/internal/league",
        "//projects/go/learn_go_with_tests/http_server/internal/player:player_lib",
    ],
)

go_binary(
    name = "http_server",
    data = ["game.db.json"],
    embed = [":http_server_lib"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "server_test",
    srcs = ["server_test.go"],
    embed = [":http_server_lib"],
    deps = [
        "//projects/go/learn_go_with_tests/http_server/internal/league",
        "//projects/go/learn_go_with_tests/http_server/internal/testutils",
    ],
)

go_test(
    name = "server_integration_test",
    srcs = ["server_integration_test.go"],
    embed = [":http_server_lib"],
    deps = [
        "//projects/go/learn_go_with_tests/http_server/internal/league",
        "//projects/go/learn_go_with_tests/http_server/internal/testutils",
    ],
)

go_test(
    name = "file_system_store_test",
    srcs = ["file_system_store_test.go"],
    embed = [":http_server_lib"],
    deps = [
        "//projects/go/learn_go_with_tests/http_server/internal/league",
        "//projects/go/learn_go_with_tests/http_server/internal/testutils",
    ],
)

go_test(
    name = "tape_test",
    srcs = ["tape_test.go"],
    embed = [":http_server_lib"],
    deps = [
        "//projects/go/learn_go_with_tests/http_server/internal/testutils",
    ],
)
