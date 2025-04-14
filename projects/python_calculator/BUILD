load("@aspect_rules_py//py:defs.bzl", "py_library", "py_test")

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
