load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "overflow",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "overflow_test",
    crate = ":overflow",
    rustc_flags = [
        "-C",
        "overflow-checks=false",
    ],
)
