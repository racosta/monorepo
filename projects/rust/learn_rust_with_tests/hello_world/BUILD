load("@rules_rust//rust:defs.bzl", "rust_binary", "rust_library", "rust_test")

rust_library(
    name = "hello_lib",
    srcs = ["src/main.rs"],
    edition = "2021",
)

rust_binary(
    name = "hello_bin",
    srcs = ["src/main.rs"],
    edition = "2021",
    deps = [":hello_lib"],
)

rust_test(
    name = "hello_test",
    crate = ":hello_lib",
)
