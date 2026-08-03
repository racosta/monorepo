load("@aspect_rules_lint//format:defs.bzl", "format_multirun")
load("@bazel_lib//lib:expand_template.bzl", "expand_template")

package(default_visibility = ["//visibility:public"])

genrule(
    name = "taplo",
    srcs = select({
        "@platforms//os:linux": ["@taplo//file"],
    }),
    outs = ["taplo_bin"],
    cmd = "gunzip -c $< > $@",
    executable = True,
)

expand_template(
    name = "taplo_wrapper",
    out = "taplo_wrapper.sh",
    data = [":taplo_bin"],
    is_executable = True,
    substitutions = {"{taplo_bin}": "$(execpath :taplo_bin)"},
    template = [
        "#!/bin/sh",
        'exec env RUST_LOG=warn "./{taplo_bin}" "$@"',
    ],
)

format_multirun(
    name = "format",
    shell = "@aspect_rules_lint//format:shfmt",
    toml = ":taplo_wrapper",
    yaml = "@aspect_rules_lint//format:yamlfmt",
)
