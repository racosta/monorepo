load("@aspect_rules_js//js:defs.bzl", "js_library")
load("@gazelle//:def.bzl", "gazelle")
load("@npm//:defs.bzl", "npm_link_all_packages")

# gazelle:prefix github.com/racosta/monorepo
# gazelle:build_file_name BUILD,BUILD.bazel
gazelle(name = "gazelle")

exports_files([".rustfmt.toml"])

npm_link_all_packages(name = "npm")

js_library(
    name = "pkg",
    srcs = ["package.json"],
    visibility = ["//visibility:public"],
)
