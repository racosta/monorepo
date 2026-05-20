"""Definition of the linter binaries

This is in its own package because it has so many loading-time symbols,
we don't want to trigger eager fetches of these for builds that don't want to run lint.
"""

load("@aspect_rules_py//py:defs.bzl", "py_binary")
load("@bazel_skylib//rules:native_binary.bzl", "native_binary")
load("@rules_python//python/entry_points:py_console_script_binary.bzl", "py_console_script_binary")

package(default_visibility = ["//:__subpackages__"])

native_binary(
    name = "clang_tidy",
    src = select(
        {
            "@bazel_tools//src/conditions:linux_x86_64": "@llvm_toolchain_llvm//:bin/clang-tidy",
            "@bazel_tools//src/conditions:linux_aarch64": "@llvm_toolchain_llvm//:bin/clang-tidy",
            "@bazel_tools//src/conditions:darwin_x86_64": "@llvm_toolchain_llvm//:bin/clang-tidy",
            "@bazel_tools//src/conditions:darwin_arm64": "@llvm_toolchain_llvm//:bin/clang-tidy",
            # llvm_toolchain doesn't support windows: https://github.com/bazel-contrib/toolchains_llvm/issues/4
            # as a workaround, you can download exes from
            # https://github.com/llvm/llvm-project/releases/tag/llvmorg-18.1.6 and make available locally.
            "@bazel_tools//src/conditions:windows_x64": "clang-tidy.exe",
        },
    ),
    out = "clang_tidy",
)

# py_console_script_binary(
#     name = "bandit",
#     pkg = "@pypi//bandit:pkg",
# )

# py_console_script_binary(
#     name = "flake8",
#     pkg = "@pypi//flake8:pkg",
# )

# py_console_script_binary(
#     name = "pylint",
#     pkg = "@pypi//pylint:pkg",
# )

py_binary(
    name = "pydoclint",
    srcs = ["pydoclint_wrapper.py"],
    main = "pydoclint_wrapper.py",
    deps = [
        "@pypi//pydoclint",
    ],
)
