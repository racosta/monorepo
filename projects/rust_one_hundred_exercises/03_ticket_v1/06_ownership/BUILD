load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "ownership",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "ownership_test",
    crate = ":ownership",
)
