load("@aspect_bazel_lib//lib:expand_template.bzl", "expand_template")
load("@aspect_rules_py//py:defs.bzl", "py_binary", "py_library", "py_test")
load("@bazel_lib//lib:testing.bzl", "assert_archive_contains")
load("@bazel_lib//lib:transitions.bzl", "platform_transition_filegroup")
load("@container_structure_test//:defs.bzl", "container_structure_test")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load")
load("//bazel:java_runtime_layer.bzl", "java_runtime_layer")
load("//bazel:py_image_layer_owned.bzl", "py_image_layer_owned")

SERVICE_NAME = "calculator_service"

PY_BINARY_NAME = "main_app"

CONTAINER_IMAGE_USER = "185"

py_library(
    name = "main_lib",
    srcs = ["main.py"],
    deps = [
        "//projects/python/calculator",
        "@pypi//flask",
    ],
)

py_binary(
    name = PY_BINARY_NAME,
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

py_image_layer_owned(
    name = "app_layer",
    binary = ":{}".format(PY_BINARY_NAME),
)

java_runtime_layer(
    name = "java_layer",
)

oci_image(
    name = "image",
    base = "@ubi10-minimal",
    entrypoint = ["/{}/{}".format(
        package_name(),
        PY_BINARY_NAME,
    )],
    env = {
        "JAVA_HOME": "/usr/lib/jvm/java-21",
        "PATH": "/usr/lib/jvm/java-21/bin:$PATH",
    },
    tars = [
        ":java_layer",
        ":app_layer_default.tar.gz",
        ":app_layer_packages.tar.gz",
        ":app_layer_interpreter.tar.gz",
    ],
    user = CONTAINER_IMAGE_USER,
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

expand_template(
    name = "stamped_img_names",
    out = "_stamped.image_names.txt",
    stamp_substitutions = {"latest": "{{BUILD_EMBED_LABEL}}"},
    template = ["gcr.io/racosta/monorepo/{}/{}:latest".format(
        package_name(),
        SERVICE_NAME,
    )],
)

oci_load(
    name = "load",
    image = ":platform_image",
    repo_tags = ":stamped_img_names",
)

assert_archive_contains(
    name = "test_java_layer",
    archive = "java_layer",
    expected = [
        "./usr/lib/jvm/java-21/bin/java",
        "./usr/lib/jvm/java-21/lib/libjava.so",
        "./usr/lib/jvm/java-21/bin/jar",
    ],
    tags = ["unit"],
    type = "tar",
)

assert_archive_contains(
    name = "test_app_layer_packages",
    archive = "app_layer_base_packages",
    expected = [
        "./{}/{}.runfiles/rules_python++pip+pypi_312_flask/site-packages/flask/__init__.py".format(
            package_name(),
            PY_BINARY_NAME,
        ),
    ],
    tags = ["unit"],
    type = "tar",
)

assert_archive_contains(
    name = "test_app_layer_interpreter",
    archive = "app_layer_base_interpreter",
    expected = [
        "./{}/{}.runfiles/rules_python++python+python_3_12_x86_64-unknown-linux-gnu/bin/python3.12".format(
            package_name(),
            PY_BINARY_NAME,
        ),
    ],
    tags = ["unit"],
    type = "tar",
)

assert_archive_contains(
    name = "test_app_layer_default",
    archive = "app_layer_base_default",
    expected = [
        "./{}/{}".format(
            package_name(),
            PY_BINARY_NAME,
        ),
        "./{}/main.py".format(
            package_name(),
        ),
    ],
    tags = ["unit"],
    type = "tar",
)

filegroup(
    name = "image_load.tar",
    srcs = [":load"],
    output_group = "tarball",
)

container_structure_test(
    name = "image_test",
    configs = ["structure_test.yaml"],
    driver = "tar",
    image = ":image_load.tar",
    tags = ["unit"],
)
