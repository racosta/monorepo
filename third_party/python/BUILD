load("@python//3.11:defs.bzl", compile_pip_requirements_3_11 = "compile_pip_requirements")
load("@python//3.12:defs.bzl", compile_pip_requirements_3_12 = "compile_pip_requirements")
load("@python//3.13:defs.bzl", compile_pip_requirements_3_13 = "compile_pip_requirements")

compile_pip_requirements_3_11(
    name = "requirements_3_11",
    src = "requirements.in",
    requirements_txt = "requirements_lock_3_11.txt",
)

compile_pip_requirements_3_12(
    name = "requirements_3_12",
    src = "requirements.in",
    requirements_txt = "requirements_lock_3_12.txt",
)

compile_pip_requirements_3_13(
    name = "requirements_3_13",
    src = "requirements.in",
    requirements_txt = "requirements_lock_3_13.txt",
)
