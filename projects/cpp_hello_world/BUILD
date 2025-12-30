load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_test")

cc_binary(
    name = "hello_world",
    srcs = ["hello_world.cpp"],
    deps = ["//projects/cpp_hello_world_lib:greeter_lib"],
)

cc_test(
    name = "hello_world_test",
    size = "small",
    srcs = ["hello_world_test.cpp"],
    deps = [
        "@googletest//:gtest",
        "@googletest//:gtest_main",
    ],
)
