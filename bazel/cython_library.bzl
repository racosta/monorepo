"""A custom Bazel macro to build Cython libraries as Python extensions."""

load("@rules_cc//cc:defs.bzl", "cc_binary")
load("@rules_python//python:defs.bzl", "py_library")

def cython_library(name, srcs, deps = [], **kwargs):
    # 1. Compile .pyx to .c
    native.genrule(
        name = name + "_cython_gen",
        srcs = srcs,
        outs = [src.replace(".pyx", ".c") for src in srcs],
        cmd = "PYTHONHASHSEED=0 $(location @cython//:cython_binary) -o $(OUTS) $(SRCS)",
        tools = ["@cython//:cython_binary"],
    )

    # 2. Compile .c into an importable native shared library extension
    cc_binary(
        name = name + ".so",
        srcs = [src.replace(".pyx", ".c") for src in srcs],
        deps = deps + ["@rules_python//python/cc:current_py_cc_headers"],
        linkshared = True,
        linkstatic = True,
        copts = ["-fvisibility=hidden"],
    )

    # 3. Expose as a standard Python library
    py_library(
        name = name,
        data = [":" + name + ".so"],
        visibility = ["//visibility:public"],
        **kwargs
    )
