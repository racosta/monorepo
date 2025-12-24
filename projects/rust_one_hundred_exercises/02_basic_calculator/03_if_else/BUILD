load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "if_else",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "if_else_test",
    crate = ":if_else",
)
