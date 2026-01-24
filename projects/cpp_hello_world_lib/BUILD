load("@rules_cc//cc:defs.bzl", "cc_library", "cc_test")

cc_library(
    name = "greeter_lib",
    srcs = ["greeter.cpp"],
    hdrs = ["greeter.h"],
    strip_include_prefix = "/" + package_name(),
    visibility = ["//visibility:public"],
)

cc_test(
    name = "greeter_test",
    size = "small",
    srcs = ["greeter_test.cpp"],
    deps = [
        ":greeter_lib",
        "@googletest//:gtest",
        "@googletest//:gtest_main",
    ],
)
