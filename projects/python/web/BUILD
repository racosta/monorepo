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
    name = "add_test",
    srcs = [
        "tests/add_test.py",
        "tests/conftest.py",
    ],
    imports = ["."],
    pytest_main = True,
    deps = [
        ":main_lib",
        "@pypi//pytest",
    ],
)
