load("@rules_tcl//tcl:tcl_library.bzl", "tcl_library")
load("@rules_tcl//tcl:tcl_test.bzl", "tcl_test")

tcl_library(
    name = "greetings",
    srcs = [
        "greet.tcl",
        "pkgIndex.tcl",
    ],
    visibility = ["//projects:__subpackages__"],
)

tcl_test(
    name = "greetings_test",
    srcs = ["greet_test.tcl"],
    deps = [":greetings"],
)
