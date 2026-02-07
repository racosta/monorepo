load("@rules_rust//rust:defs.bzl", "rust_binary", "rust_library", "rust_test")

rust_library(
    name = "iteration_lib",
    srcs = ["src/main.rs"],
    edition = "2021",
)

rust_binary(
    name = "iteration_bin",
    srcs = ["src/main.rs"],
    edition = "2021",
    deps = [":iteration_lib"],
)

rust_test(
    name = "iteration_test",
    crate = ":iteration_lib",
)
