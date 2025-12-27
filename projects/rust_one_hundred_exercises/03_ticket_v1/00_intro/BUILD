load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "intro_02",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "intro_02_test",
    crate = ":intro_02",
)
