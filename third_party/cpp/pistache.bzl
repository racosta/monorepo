"""This file defines an extension for the Pistache C++ HTTP library."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _pistache_repo_impl(_module_ctx):
    http_archive(
        name = "pistache",
        urls = ["https://github.com/pistacheio/pistache/archive/refs/tags/v0.4.26.tar.gz"],
        sha256 = "29af6562547497acf6f49170661786fe8cf1ed3712ad80e69c53da4661c59544",
        strip_prefix = "pistache-0.4.26",
        build_file_content = """
load("@rules_cc//cc:defs.bzl", "cc_library")

cc_library(
    name = "pistache",
    srcs = glob(["src/**/*.cc"]),
    hdrs = glob(["include/pistache/*.h", "include/pistache/**/*.h"]),
    includes = ["include"],
    visibility = ["//visibility:public"],
    linkopts = ["-lpthread"],
    deps = [
        # You may need to add OpenSSL here if you need HTTPS
        "@boringssl//:ssl",
        "@howardhinnant_date//:howardhinnant_date",
    ],
)
""",
    )

pistache_extension = module_extension(implementation = _pistache_repo_impl)
