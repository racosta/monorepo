load("@rules_java//java:defs.bzl", "java_library", "java_test")

java_library(
    name = "greeter",
    srcs = ["src/main/java/com/github/racosta/greeter/Greeter.java"],
    visibility = ["//visibility:public"],
)

java_test(
    name = "greeter_test",
    srcs = ["src/test/java/com/github/racosta/greeter/GreeterTest.java"],
    test_class = "com.github.racosta.greeter.GreeterTest",
    deps = [":greeter"],
)
