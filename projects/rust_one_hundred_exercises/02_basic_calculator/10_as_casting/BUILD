load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "as_cast",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "as_cast_test",
    crate = ":as_cast",
)
