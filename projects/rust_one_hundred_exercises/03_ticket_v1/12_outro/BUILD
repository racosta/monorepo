load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

rust_library(
    name = "outro_02",
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "outro_02_integration_test",
    srcs = ["tests/integration.rs"],
    deps = [":outro_02"],
)
