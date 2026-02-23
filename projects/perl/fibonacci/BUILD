load("@rules_perl//perl:perl.bzl", "perl_library", "perl_test")

perl_library(
    name = "fibonacci",
    srcs = ["fibonacci.pm"],
)

perl_test(
    name = "fibonacci_test",
    srcs = ["fibonacci_test.t"],
    deps = [":fibonacci"],
)
