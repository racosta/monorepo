load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "for",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "for_test",
    crate = ":for",
)
