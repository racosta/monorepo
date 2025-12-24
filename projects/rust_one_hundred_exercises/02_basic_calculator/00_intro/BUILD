load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "intro_01",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "intro_01_test",
    crate = ":intro_01",
)
