load("@aspect_rules_py//py:defs.bzl", "py_library", "py_test")
load("//tools/lint:linters.bzl", "pydoclint_test", "ruff_test", "ty_test")

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
        "@pypi//pytest_cov",
    ],
)

py_library(
    name = "test_srcs",
    srcs = glob(["tests/**/*.py"]),
    visibility = ["//projects/python/calculator:__subpackages__"],
)

pydoclint_test(
    name = "pydoclint_test",
    srcs = [
        ":calculator",
        ":test_srcs",
    ],
)

ruff_test(
    name = "ruff_test",
    srcs = [
        ":calculator",
        ":test_srcs",
    ],
)

ty_test(
    name = "ty_test",
    srcs = [
        ":calculator",
        ":calculator_test",
    ],
)
