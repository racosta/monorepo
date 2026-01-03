load("@rules_rust//rust:defs.bzl", "rust_library", "rust_test")

LIB_NAME = "outro_02"

rust_library(
    name = LIB_NAME,
    srcs = ["src/lib.rs"],
    edition = "2021",
)

rust_test(
    name = "{lib_name}_integration_test".format(lib_name = LIB_NAME),
    srcs = ["tests/integration.rs"],
    deps = [":{lib_name}".format(lib_name = LIB_NAME)],
)
