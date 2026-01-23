load("@aspect_rules_js//js:defs.bzl", "js_binary", "js_library")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "npm")

js_library(
    name = "tsc",
    srcs = ["src/hello.mjs"],
    deps = [
        ":npm/@monorepo/js_hello_world_lib",
    ],
)

js_binary(
    name = "hello_world",
    data = [
        ":tsc",
    ],
    entry_point = "src/hello.mjs",
)
