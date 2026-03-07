"""This file defines the repositories for non-Bazel C++ dependencies used in the project."""

def _single_header_repo_impl(repository_ctx):
    repository_ctx.download(
        url = repository_ctx.attr.url,
        output = "ApprovalTests.hpp",
        sha256 = repository_ctx.attr.sha256,
    )

    repository_ctx.file("BUILD", """
load("@rules_cc//cc:defs.bzl", "cc_library")

cc_library(
    name = "approvaltests",
    hdrs = ["ApprovalTests.hpp"],
    visibility = ["//visibility:public"],
)
""")

single_header_repo = repository_rule(
    implementation = _single_header_repo_impl,
    attrs = {
        "url": attr.string(mandatory = True),
        "sha256": attr.string(),
    },
)

def _approval_extension_impl(_module_ctx):
    single_header_repo(
        name = "approvaltests_cpp",
        url = "https://github.com/approvals/ApprovalTests.cpp/releases/download/v.10.13.0/ApprovalTests.v.10.13.0.hpp",
        sha256 = "c00f6390b81d9924dc646e9d32b61e1e09abda106c13704f714ac349241bb9ff",
    )

approval_extension = module_extension(implementation = _approval_extension_impl)
