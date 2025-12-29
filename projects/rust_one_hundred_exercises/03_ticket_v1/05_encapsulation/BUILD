load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "encapsulation",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "encapsulation_test",
    crate = ":encapsulation",
)
