load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "modules",
    srcs = ["src/lib.rs"],
    edition = "2021",
    deps = ["//projects/rust_one_hundred_exercises/helpers/common"],
)

rust_test(
    name = "modules_test",
    crate = ":modules",
)
