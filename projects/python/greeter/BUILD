load("@aspect_rules_py//py:defs.bzl", "py_library", "py_test")
load("//tools/lint:linters.bzl", "pydoclint_test", "ruff_test", "ty_test")

py_library(
    name = "greeter",
    srcs = [
        "__init__.py",
        "greeter.py",
    ],
    imports = ["."],
    visibility = ["//visibility:public"],
)

py_test(
    name = "greeter_test",
    srcs = ["tests/greeter_test.py"],
    pytest_main = True,
    deps = [
        ":greeter",
        "@pypi//pytest",
        "@pypi//pytest_cov",
    ],
)

py_library(
    name = "test_srcs",
    srcs = glob(["tests/**/*.py"]),
    visibility = ["//projects/python/greeter:__subpackages__"],
    deps = [
        ":greeter",
        "@pypi//pytest",
    ],
)

pydoclint_test(
    name = "pydoclint_test",
    srcs = [
        ":greeter",
        ":test_srcs",
    ],
)

ruff_test(
    name = "ruff_test",
    srcs = [
        ":greeter",
        ":test_srcs",
    ],
)

ty_test(
    name = "ty_test",
    srcs = [
        ":greeter",
        ":greeter_test",
    ],
)
