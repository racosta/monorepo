load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "json_generator_lib",
    srcs = ["json_generator.cpp"],
    hdrs = ["json_generator.h"],
    deps = ["@rapidjson"],
)

cc_binary(
    name = "rapidjson",
    srcs = ["main.cpp"],
    deps = [
        ":json_generator_lib",
        "@spdlog",
    ],
)

cc_test(
    name = "json_generator_test",
    size = "small",
    srcs = ["json_generator_test.cpp"],
    deps = [
        ":json_generator_lib",
        "@googletest//:gtest",
        "@googletest//:gtest_main",
    ],
)
