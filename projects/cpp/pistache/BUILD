load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "hello_server",
    srcs = ["main.cpp"],
    deps = [
        "//projects/cpp/fmt_adapters:pistache",
        "@pistache",
        "@spdlog",
    ],
)
