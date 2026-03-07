load("@rules_cc//cc:defs.bzl", "cc_test")

cc_test(
    name = "googletest_starter_test",
    size = "small",
    srcs = [
        "googletest_starter_main.cpp",
        "googletest_starter_test.cpp",
    ],
    data = glob(["approval_tests/**"]),
    deps = [
        "@approvaltests_cpp//:approvaltests",
        "@googletest//:gtest",
        "@googletest//:gtest_main",
    ],
)
