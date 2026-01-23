load("@aspect_rules_jest//jest:defs.bzl", "jest_test")
load("@aspect_rules_js//js:defs.bzl", "js_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "npm")

js_library(
    name = "tsc",
    srcs = ["greeter.mjs"],
)

js_library(
    name = "pkg",
    srcs = ["package.json"],
    visibility = ["//visibility:public"],
    deps = [":tsc"],
)

jest_test(
    name = "greeter_test",
    config = "jest.config.mjs",
    data = [
        "greeter.test.mjs",
        "package.json",
        ":tsc",
    ],
    node_modules = ":npm",
    node_options = [
        "--experimental-vm-modules",
    ],
)
