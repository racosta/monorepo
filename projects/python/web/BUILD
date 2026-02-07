load("@aspect_rules_py//py:defs.bzl", "py_binary", "py_library", "py_test")

py_library(
    name = "main_lib",
    srcs = ["main.py"],
    deps = [
        "//projects/python/calculator",
        "@pypi//flask",
    ],
)

py_binary(
    name = "main",
    srcs = ["main.py"],
    deps = [
        ":main_lib",
    ],
)

py_test(
    name = "test_add",
    srcs = [
        "tests/conftest.py",
        "tests/test_add.py",
    ],
    imports = ["."],
    deps = [
        ":main_lib",
        "@pypi//pytest",
    ],
)
