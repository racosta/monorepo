load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "destructor",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "destructor_test",
    crate = ":destructor",
)
