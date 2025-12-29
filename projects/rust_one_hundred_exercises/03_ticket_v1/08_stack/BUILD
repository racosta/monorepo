load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "stack",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "stack_test",
    crate = ":stack",
)
