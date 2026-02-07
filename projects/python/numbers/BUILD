load("@aspect_rules_py//py:defs.bzl", "py_library", "py_test")

py_library(
    name = "numbers",
    srcs = [
        "is_leap_year.py",
    ],
)

py_test(
    name = "test_is_leap_year",
    srcs = ["tests/test_is_leap_year.py"],
    imports = ["."],
    deps = [
        ":numbers",
        "@pypi//pytest",
    ],
)
