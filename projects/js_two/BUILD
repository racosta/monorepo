load("@aspect_rules_js//js:defs.bzl", "js_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "npm")

js_library(
    name = "tsc",
    srcs = ["src/lib.js"],
    deps = [
        ":npm/@monorepo/js_first",
        ":npm/cowsay",
    ],
)

js_library(
    name = "pkg",
    srcs = ["package.json"],
    visibility = ["//visibility:public"],
    deps = [":tsc"],
)
