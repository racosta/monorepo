"""A custom Bazel macro to build Cython libraries as Python extensions."""

load("@rules_cc//cc:defs.bzl", "cc_binary")
load("@rules_python//python:defs.bzl", "py_library")

def cython_library(name, srcs, deps = [], **kwargs):
    """A macro to build a Cython library as a Python extension.

    Args:
        name: The name of the target.
        srcs: A list of .pyx source files to compile.
        deps: A list of dependencies for the Cython library.
        **kwargs: Additional keyword arguments to pass to the py_library rule.
    """

    # Separate .pyx and .pyi files if passed in srcs
    pyx_srcs = [s for s in srcs if s.endswith(".pyx")]
    pyi_srcs = [s for s in srcs if s.endswith(".pyi")]
    py_srcs = [s for s in srcs if s.endswith(".py")]

    if not pyx_srcs:
        fail("cython_library '%s' must contain at least one .pyx file in srcs" % name)

    c_srcs = [s.replace(".pyx", ".c") for s in pyx_srcs]

    # 1. Compile .pyx files to .c
    native.genrule(
        name = name + "_cython_gen",
        srcs = pyx_srcs,
        outs = c_srcs,
        cmd = "PYTHONHASHSEED=0 $(location @cython//:cython_binary) -o $(OUTS) " + " ".join(["$(location %s)" % s for s in pyx_srcs]),
        tools = ["@cython//:cython_binary"],
    )

    # 2. Compile .c into an importable native shared library extension
    cc_binary(
        name = name + ".so",
        srcs = c_srcs,
        deps = deps + ["@rules_python//python/cc:current_py_cc_headers"],
        linkshared = True,
        linkstatic = True,
        copts = ["-fvisibility=hidden"],
    )

    # 3. Copy .pyi to .py (same base name) so Pyrefly registers `math_utils` module
    py_stub_targets = []
    for pyi in pyi_srcs:
        stub_py_name = pyi.replace(".pyi", ".py")  # Keep exact module name!
        native.genrule(
            name = name + "_" + pyi.replace(".", "_") + "_gen_py",
            srcs = [pyi],
            outs = [stub_py_name],
            cmd = "cp $< $@",
        )
        py_stub_targets.append(stub_py_name)

    # 4. Expose as a standard Python library
    py_library(
        name = name,
        srcs = py_srcs + py_stub_targets,
        data = [":" + name + ".so"] + pyi_srcs,
        visibility = ["//visibility:public"],
        **kwargs
    )
