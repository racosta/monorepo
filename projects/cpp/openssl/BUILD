load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "ssl_checker_lib",
    srcs = ["ssl_checker.cpp"],
    hdrs = ["ssl_checker.h"],
    deps = [
        "@boringssl//:ssl",
    ],
)

cc_binary(
    name = "openssl_example",
    srcs = ["main.cpp"],
    deps = [
        ":ssl_checker_lib",
        "@boringssl//:ssl",
        "@spdlog",
    ],
)

cc_test(
    name = "ssl_checker_test",
    size = "small",
    srcs = ["ssl_checker_test.cpp"],
    deps = [
        ":ssl_checker_lib",
        "@googletest//:gtest",
        "@googletest//:gtest_main",
    ],
)
