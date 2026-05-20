"""Definition of the linter binaries

This is in its own package because it has so many loading-time symbols,
we don't want to trigger eager fetches of these for builds that don't want to run lint.
"""

load("@aspect_rules_py//py:defs.bzl", "py_binary")
load("@rules_python//python/entry_points:py_console_script_binary.bzl", "py_console_script_binary")

package(default_visibility = ["//:__subpackages__"])

py_console_script_binary(
    name = "bandit",
    pkg = "@pypi//bandit:pkg",
)

py_console_script_binary(
    name = "flake8",
    pkg = "@pypi//flake8:pkg",
)

py_console_script_binary(
    name = "pylint",
    pkg = "@pypi//pylint:pkg",
)

py_binary(
    name = "pydoclint",
    srcs = ["pydoclint_wrapper.py"],
    main = "pydoclint_wrapper.py",
    deps = [
        "@pypi//pydoclint",
    ],
)
