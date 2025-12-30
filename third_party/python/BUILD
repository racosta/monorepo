load("@rules_uv//uv:pip.bzl", "pip_compile")

pip_compile(
    name = "requirements_3_11",
    py3_runtime = "@python_3_11//:py3_runtime",
    requirements_in = "requirements.in",
    requirements_txt = "requirements_lock_3_11.txt",
)

pip_compile(
    name = "requirements_3_12",
    py3_runtime = "@python_3_12//:py3_runtime",
    requirements_in = "requirements.in",
    requirements_txt = "requirements_lock_3_12.txt",
)

pip_compile(
    name = "requirements_3_13",
    py3_runtime = "@python_3_13//:py3_runtime",
    requirements_in = "requirements.in",
    requirements_txt = "requirements_lock_3_13.txt",
)
