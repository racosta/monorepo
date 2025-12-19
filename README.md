# monorepo

Starting with a Nix shell and pre-commit hooks, then based on [kriscfoster/multi-language-bazel-monorepo](https://github.com/kriscfoster/multi-language-bazel-monorepo).

Investigating how new features work with Bazel and bzlmod, etc.

## Devenv

In this repository, the environment is managed using [devenv](devenv.sha) and includes Bazel and other tools used in
this repository.

Nix is a prerequisite. Visit [devenv::Getting started](https://devenv.sh/getting-started/) and use the
instructions there. After cloning, start the devenv shell by changing into the repo root directory
and using `devenv shell`.

## Pre-commit

This repository also makes use of the [pre-commit](https://pre-commit.com/) framework. The pre-commit
tool itself is included in the Nix shell, but the hooks must be installed after cloning using the
command `pre-commit install`.
