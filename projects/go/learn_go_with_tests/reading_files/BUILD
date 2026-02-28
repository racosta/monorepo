load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "reading_files",
    srcs = [
        "blogposts.go",
        "post.go",
    ],
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/reading_files",
    visibility = ["//visibility:public"],
)

go_test(
    name = "reading_files_test",
    srcs = ["blogposts_test.go"],
    deps = [":reading_files"],
)
