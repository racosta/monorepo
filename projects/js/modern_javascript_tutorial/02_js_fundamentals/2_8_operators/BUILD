load("@aspect_rules_jest//jest:defs.bzl", "jest_test")
load("@aspect_rules_js//js:defs.bzl", "js_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages()

js_library(
    name = "operators_lib",
    srcs = ["src/operators.mjs"],
)

jest_test(
    name = "operators_test",
    config = "tests/jest.config.mjs",
    data = [
        "package.json",
        "tests/operators.test.mjs",
        ":operators_lib",
    ],
    node_modules = ":node_modules",
    node_options = [
        "--experimental-vm-modules",
    ],
)
