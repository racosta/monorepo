load("@rules_rust//rust:defs.bzl", "rust_binary")

rust_binary(
    name = "rust_hello_world",
    srcs = ["src/main.rs"],
    deps = [
        "//projects/rust_hello_world_lib:hello_lib",
        "@crates//:anyhow",
    ],
)
