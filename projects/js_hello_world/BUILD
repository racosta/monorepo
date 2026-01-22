load("@aspect_rules_js//js:defs.bzl", "js_binary")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "node_modules")

js_binary(
    name = "hello_world",
    entry_point = "src/hello.mjs",
    visibility = ["//visibility:public"],
)
