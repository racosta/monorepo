load("@rules_cc//cc:defs.bzl", "cc_binary")
# load("//tools/lint:linters.bzl", "clang_tidy_test")

cc_binary(
    name = "hello_server",
    srcs = ["main.cpp"],
    deps = [
        "//projects/cpp/fmt_adapters:pistache",
        "@pistache",
        "@spdlog",
    ],
)

# clang_tidy_test(
#     name = "clang_tidy",
#     srcs = [":hello_server"],
# )
