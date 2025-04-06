load("@rules_python//python:defs.bzl", "py_test")
load("@rules_python//python:py_library.bzl", "py_library")

py_library(
    name = "calculator",
    srcs = ["calculator.py"],
    imports = ["."],
    visibility = ["//projects/python_web:__pkg__"],
    deps = [],
)

py_test(
    name = "calculator_test",
    srcs = ["calculator_test.py"],
    deps = [
        ":calculator",
    ],
)
