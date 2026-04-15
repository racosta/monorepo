load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "openssl",
    srcs = ["main.cpp"],
    deps = [
        "@boringssl//:ssl",
        "@spdlog",
    ],
)
