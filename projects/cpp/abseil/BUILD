load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "abseil_example_lib",
    srcs = ["abseil_example.cpp"],
    hdrs = ["abseil_example.h"],
    deps = ["@abseil-cpp//absl/strings"],
)

cc_binary(
    name = "abseil_example",
    srcs = ["main.cpp"],
    deps = [
        ":abseil_example_lib",
        "@spdlog",
    ],
)

cc_test(
    name = "abseil_example_test",
    size = "small",
    srcs = ["abseil_example_test.cpp"],
    deps = [
        ":abseil_example_lib",
        "@googletest//:gtest",
        "@googletest//:gtest_main",
    ],
)
