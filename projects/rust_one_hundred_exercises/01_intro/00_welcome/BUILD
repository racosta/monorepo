load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "welcome_00",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "welcome_00_test",
    crate = ":welcome_00",
)
