# monorepo

## Nix

In this repository, a Nix shell configuration is included in `shell.nix` which includes Bazel and
other development tools for this repository.

Nix is a prerequisite. Visit [Install Nix](https://nix.dev/install-nix.html) and use the
instructions there. After cloning, start the Nix shell by changing into the repo root directory
and using `nix-shell`.

## Pre-commit

This repository also makes use of the [pre-commit](https://pre-commit.com/) framework. The pre-commit
tool itself is included in the Nix shell, but the hooks must be installed after cloning using the
command `pre-commit install`.
