load("@aspect_rules_py//py:defs.bzl", "py_binary", "py_library")
load("//tools/lint:linters.bzl", "pydoclint_test", "ruff_test", "ty_test")

py_library(
    name = "hello_world",
    srcs = ["hello_world.py"],
    imports = ["."],
    deps = [
        "//projects/python/greeter",
    ],
)

py_binary(
    name = "hello_world_app",
    srcs = ["hello_world.py"],
    deps = [":hello_world"],
)

pydoclint_test(
    name = "pydoclint_test",
    srcs = [
        ":hello_world",
    ],
)

ruff_test(
    name = "ruff_test",
    srcs = [
        ":hello_world",
    ],
)

ty_test(
    name = "ty_test",
    srcs = [
        ":hello_world",
    ],
)
