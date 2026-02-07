load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "greeter",
    srcs = [
        "src/greeter.rs",
        "src/lib.rs",
    ],
    visibility = ["//projects/rust/hello_world:__pkg__"],
    deps = [
        "@crates//:anyhow",
    ],
)

rust_test(
    name = "greeter_test",
    crate = ":greeter",
    deps = [
        "@crates//:anyhow",
    ],
)
