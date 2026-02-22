load("@aspect_rules_jest//jest:defs.bzl", "jest_test")
load("@aspect_rules_js//js:defs.bzl", "js_binary", "js_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "npm")

js_library(
    name = "string_utils",
    srcs = ["src/stringUtils.mjs"],
)

js_binary(
    name = "data_types",
    data = [":string_utils"],
    entry_point = "src/interpolation.mjs",
)

jest_test(
    name = "data_types_test",
    config = "tests/jest.config.mjs",
    data = [
        "package.json",
        "tests/string_utils.test.mjs",
        ":string_utils",
    ],
    node_modules = ":npm",
    node_options = [
        "--experimental-vm-modules",
    ],
)
