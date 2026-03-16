load("@aspect_rules_js//js:defs.bzl", "js_binary", "js_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages()

js_library(
    name = "tsc",
    srcs = ["src/hello.mjs"],
    deps = [
        ":node_modules/@monorepo/greeter",
    ],
)

js_binary(
    name = "hello_world",
    data = [
        ":tsc",
    ],
    entry_point = "src/hello.mjs",
)
