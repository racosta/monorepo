"""Define Python linter aspects"""

load("@aspect_rules_lint//lint:bandit.bzl", "lint_bandit_aspect")
load("@aspect_rules_lint//lint:buildifier.bzl", "lint_buildifier_aspect")
load("@aspect_rules_lint//lint:clang_tidy.bzl", "lint_clang_tidy_aspect")
load("@aspect_rules_lint//lint:lint_test.bzl", "lint_test")
load("@aspect_rules_lint//lint:pydoclint.bzl", "lint_pydoclint_aspect")
load("@aspect_rules_lint//lint:ruff.bzl", "lint_ruff_aspect")
load("@aspect_rules_lint//lint:shellcheck.bzl", "lint_shellcheck_aspect")
load("@aspect_rules_lint//lint:ty.bzl", "lint_ty_aspect")
load("@aspect_rules_lint//lint:yamllint.bzl", "lint_yamllint_aspect")

bandit = lint_bandit_aspect(
    binary = Label("//tools/lint:bandit"),
    config = Label("//:pyproject.toml"),
)

bandit_test = lint_test(aspect = bandit)

buildifier = lint_buildifier_aspect(
    binary = Label("@buildifier_prebuilt//:buildifier"),
)

buildifier_test = lint_test(aspect = buildifier)

clang_tidy = lint_clang_tidy_aspect(
    binary = Label("//tools/lint:clang_tidy"),
    configs = [
        Label("//:.clang-tidy"),
    ],
    lint_target_headers = True,
    angle_includes_are_system = False,
    verbose = False,
)

clang_tidy_test = lint_test(aspect = clang_tidy)

pydoclint = lint_pydoclint_aspect(
    binary = Label("//tools/lint:pydoclint"),
    config = Label("//:pyproject.toml"),
)

pydoclint_test = lint_test(aspect = pydoclint)

ruff = lint_ruff_aspect(
    binary = Label("@aspect_rules_lint//lint:ruff_bin"),
    configs = [
        Label("@//:pyproject.toml"),
    ],
)

ruff_test = lint_test(aspect = ruff)

shellcheck = lint_shellcheck_aspect(
    binary = Label("@aspect_rules_lint//lint:shellcheck_bin"),
    config = Label("@//:.shellcheckrc"),
)

shellcheck_test = lint_test(aspect = shellcheck)

ty = lint_ty_aspect(
    binary = Label("@aspect_rules_lint//lint:ty_bin"),
    config = Label("@//:pyproject.toml"),
)

ty_test = lint_test(aspect = ty)

yamllint = lint_yamllint_aspect(
    binary = Label("//tools/lint:yamllint"),
    config = Label("//:.yamllint.yaml"),
    extra_args = ["--strict"],
)

yamllint_test = lint_test(aspect = yamllint)
