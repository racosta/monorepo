load("@aspect_rules_py//py:defs.bzl", "py_library", "py_test")
load("@pypi//:requirements.bzl", "requirement")

py_library(
    name = "calculator",
    srcs = [
        "__init__.py",
        "calculator.py",
    ],
    imports = ["."],
    visibility = ["//projects/python_web:__pkg__"],
    deps = [],
)

py_test(
    name = "calculator_test",
    srcs = ["calculator_test.py"],
    deps = [
        ":calculator",
        requirement("pytest"),
    ],
)
