load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "nlohmann_json_lib",
    srcs = ["json_generator.cpp"],
    hdrs = ["json_generator.h"],
    deps = [
        "@nlohmann_json//:json",
    ],
)

cc_binary(
    name = "nlohmann_json",
    srcs = ["main.cpp"],
    deps = [
        ":nlohmann_json_lib",
        "@nlohmann_json//:json",
        "@spdlog",
    ],
)

cc_test(
    name = "json_generator_test",
    size = "small",
    srcs = ["json_generator_test.cpp"],
    deps = [
        ":nlohmann_json_lib",
        "@googletest//:gtest",
        "@googletest//:gtest_main",
    ],
)
