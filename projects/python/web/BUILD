load("@aspect_rules_py//py:defs.bzl", "py_binary", "py_image_layer", "py_library", "py_test")
load("@bazel_lib//lib:transitions.bzl", "platform_transition_filegroup")
load("@container_structure_test//:defs.bzl", "container_structure_test")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load")

py_library(
    name = "main_lib",
    srcs = ["main.py"],
    deps = [
        "//projects/python/calculator",
        "@pypi//flask",
    ],
)

py_binary(
    name = "main_app",
    srcs = ["main.py"],
    deps = [
        ":main_lib",
    ],
)

py_test(
    name = "add_test",
    srcs = [
        "tests/add_test.py",
        "tests/conftest.py",
    ],
    imports = ["."],
    pytest_main = True,
    deps = [
        ":main_lib",
        "@pypi//pytest",
    ],
)

py_image_layer(
    name = "app_layer",
    binary = ":main_app",
)

oci_image(
    name = "image",
    base = "@ubuntu",
    entrypoint = ["/{}/main_app".format(package_name())],
    tars = [":app_layer"],
)

platform(
    name = "x86_64_linux",
    constraint_values = [
        "@platforms//os:linux",
        "@platforms//cpu:x86_64",
    ],
)

platform_transition_filegroup(
    name = "platform_image",
    srcs = [":image"],
    target_platform = select({
        "@platforms//cpu:arm64": ":aarch64_linux",
        "@platforms//cpu:x86_64": ":x86_64_linux",
    }),
)

oci_load(
    name = "image_load",
    image = ":platform_image",
    repo_tags = ["gcr.io/oci_python_hello_world:latest"],
)

filegroup(
    name = "image_load.tar",
    srcs = [":image_load"],
    output_group = "tarball",
)

container_structure_test(
    name = "image_test",
    configs = ["structure_test.yaml"],
    driver = "tar",
    image = ":image_load.tar",
    tags = ["unit"],
)
