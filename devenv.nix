{ pkgs, ... }:

{
  # Packages available in the development shell
  packages = with pkgs; [
    bat
    bazel_8
    bazel-buildtools
    difftastic
    dua
    eza
    git
    go
    golangci-lint
    jq
    just
    kondo
    lazygit
    lcov
    mdcat
    openssl
    pre-commit
    python3
    readline
    ripgrep
    starship
    tokei
    uv
    vim
    zlib
  ];

  # Environment variables
  env.GREETING = "Hello, Nix!";

  # # Enable devenv's built‑in pre-commit integration
  # pre-commit = {
  #   enable = true;

  #   # devenv will automatically install hooks from .pre-commit-config.yaml
  #   # and run them on `git commit`.
  #   hooks = {
  #     trailing-whitespace.enable = true;
  #     end-of-file-fixer.enable = true;
  #   };
  # };

  # Shell hook (runs when entering `devenv shell`)
  enterShell = ''
    alias ls='eza --icons'
    eval "$(starship init bash)"
  '';
}
