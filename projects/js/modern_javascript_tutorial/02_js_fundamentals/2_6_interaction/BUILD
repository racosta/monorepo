load("@aspect_rules_jest//jest:defs.bzl", "jest_test")
load("@aspect_rules_js//js:defs.bzl", "js_binary", "js_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "npm")

js_library(
    name = "interaction_lib",
    srcs = ["src/greeter.mjs"],
)

js_binary(
    name = "interaction",
    data = [":interaction_lib"],
    entry_point = "src/interaction.mjs",
)

jest_test(
    name = "interaction_test",
    config = "tests/jest.config.mjs",
    data = [
        "package.json",
        "tests/greeter.test.mjs",
        ":interaction_lib",
    ],
    node_modules = ":npm",
    node_options = [
        "--experimental-vm-modules",
    ],
)
