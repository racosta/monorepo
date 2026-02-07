load("@rules_rust//rust:defs.bzl", "rust_binary", "rust_library", "rust_test")

rust_library(
    name = "integers_lib",
    srcs = ["src/main.rs"],
    edition = "2021",
)

rust_binary(
    name = "integers_bin",
    srcs = ["src/main.rs"],
    edition = "2021",
    deps = [":integers_lib"],
)

rust_test(
    name = "integers_test",
    crate = ":integers_lib",
)
