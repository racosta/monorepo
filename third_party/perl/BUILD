load("@rules_perl//perl:perl.bzl", "perl_binary")

perl_binary(
    name = "cover",
    srcs = ["cover.pl"],
    main = "cover.pl",
    visibility = ["//visibility:public"],
    deps = [
        "@cpan_deps//Devel-Cover",
    ],
)
