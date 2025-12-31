load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "references_in_memory",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "references_in_memory_test",
    crate = ":references_in_memory",
)
