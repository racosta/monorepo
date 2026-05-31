load("@aspect_rules_py//py:defs.bzl", "py_binary", "py_library", "py_test")
load("//bazel:cython_library.bzl", "cython_library")
load("//tools/lint:linters.bzl", "pydoclint_test", "ruff_test", "ty_test")

cython_library(
    name = "math_utils",
    srcs = ["math_utils.pyx"],
)

py_binary(
    name = "math_utils_app",
    srcs = ["main.py"],
    deps = [":math_utils"],
)

py_test(
    name = "math_utils_test",
    srcs = ["tests/math_utils_test.py"],
    pytest_main = True,
    deps = [
        ":math_utils",
        "@pypi//pytest",
    ],
)

py_library(
    name = "test_srcs",
    srcs = glob(["tests/**/*.py"]),
    visibility = ["//projects/python/cython_example:__subpackages__"],
)

pydoclint_test(
    name = "pydoclint_test",
    srcs = [
        ":math_utils",
        ":test_srcs",
    ],
)

ruff_test(
    name = "ruff_test",
    srcs = [
        ":math_utils",
        ":test_srcs",
    ],
)

ty_test(
    name = "ty_test",
    srcs = [
        ":math_utils",
        ":math_utils_test",
    ],
)
