load("@aspect_rules_py//py:defs.bzl", "py_library", "py_test")
load("//tools/lint:linters.bzl", "bandit_test", "pydoclint_test", "ruff_test")

py_library(
    name = "numbers",
    srcs = [
        "__init__.py",
        "is_leap_year.py",
    ],
)

py_test(
    name = "is_leap_year_test",
    srcs = ["tests/is_leap_year_test.py"],
    pytest_main = True,
    deps = [
        ":numbers",
        "@pypi//pytest",
    ],
)

py_library(
    name = "test_srcs",
    srcs = glob(["tests/**/*.py"]),
    visibility = ["//projects/python/numbers:__subpackages__"],
)

pydoclint_test(
    name = "pydoclint_test",
    srcs = [
        ":numbers",
        ":test_srcs",
    ],
)

ruff_test(
    name = "ruff_test",
    srcs = [
        ":numbers",
        ":test_srcs",
    ],
)
