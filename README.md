# monorepo

Starting with a Nix shell and pre-commit hooks, then based on [kriscfoster/multi-language-bazel-monorepo](https://github.com/kriscfoster/multi-language-bazel-monorepo).

Investigating how new features work with Bazel and bzlmod, etc.

## Devenv

In this repository, the environment is managed using [devenv](devenv.sh) and includes Bazel and other tools used in
this repository.

Nix is a prerequisite. Visit [devenv::Getting started](https://devenv.sh/getting-started/) and use the
instructions there. After cloning, start the devenv shell by changing into the repo root directory
and using `devenv shell`.

## Pre-commit

Pre-commit hooks are managed in devenv via [Git hooks](https://devenv.sh/git-hooks/).
