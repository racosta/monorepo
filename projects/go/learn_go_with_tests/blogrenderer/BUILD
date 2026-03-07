load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "blogrenderer",
    srcs = [
        "post.go",
        "post_view_model.go",
        "renderer.go",
    ],
    embedsrcs = glob(["templates/*.gohtml"]),
    importpath = "github.com/racosta/monorepo/projects/go/learn_go_with_tests/blogrenderer",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_gomarkdown_markdown//:markdown",
        "@com_github_gomarkdown_markdown//parser",
    ],
)

go_test(
    name = "blogrenderer_test",
    srcs = ["renderer_test.go"],
    data = glob(["testdata/*"]),
    deps = [
        ":blogrenderer",
        "@com_github_sebdah_goldie_v2//:goldie",
    ],
)
