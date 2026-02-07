load("@aspect_rules_py//py:defs.bzl", "py_library", "py_test")

py_library(
    name = "calculator",
    srcs = [
        "__init__.py",
        "calculator.py",
    ],
    imports = ["."],
    visibility = ["//projects/python/web:__pkg__"],
    deps = [],
)

py_test(
    name = "calculator_test",
    srcs = ["tests/calculator_test.py"],
    pytest_main = True,
    deps = [
        ":calculator",
        "@pypi//pytest",
    ],
)
