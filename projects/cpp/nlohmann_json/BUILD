load("@rules_cc//cc:defs.bzl", "cc_binary")

cc_binary(
    name = "nlohmann_json",
    srcs = ["main.cpp"],
    cxxopts = ["-std=c++17"],
    deps = [
        "@nlohmann_json//:json",
        "@spdlog",
    ],
)
