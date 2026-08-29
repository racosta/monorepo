load("@aspect_rules_py//py:defs.bzl", "py_binary", "py_library", "py_test")
load("//tools/lint:linters.bzl", "pydoclint_test", "ruff_test", "ty_test")

py_library(
    name = "fastapi_lib",
    srcs = ["main.py"],
    deps = [
        "@pypi//fastapi",
        "@pypi//uvicorn",
    ],
)

py_binary(
    name = "fastapi_app",
    srcs = ["main.py"],
    deps = [
        ":fastapi_lib",
    ],
)

py_test(
    name = "test_read_root",
    srcs = [
        "tests/conftest.py",
        "tests/test_read_root.py",
    ],
    imports = ["."],
    pytest_main = True,
    deps = [
        ":fastapi_lib",
        "@pypi//httpx2",
        "@pypi//pytest",
        "@pypi//pytest_cov",
    ],
)

py_library(
    name = "test_srcs",
    srcs = glob(["tests/**/*.py"]),
    visibility = ["//projects/python/fastapi:__subpackages__"],
)

pydoclint_test(
    name = "pydoclint_test",
    srcs = [
        ":fastapi_lib",
        ":test_srcs",
    ],
)

ruff_test(
    name = "ruff_test",
    srcs = [
        ":fastapi_lib",
        ":test_srcs",
    ],
)

ty_test(
    name = "ty_test",
    srcs = [
        ":fastapi_lib",
        ":test_read_root",
    ],
)
