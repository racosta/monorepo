{ pkgs, lib, config, ... }:

{
  # 1. Define the "base" configuration (optional - common to all)
  # Packages used everywhere (Dev + CI)
  packages = with pkgs; [
    bazel_8
    bazel-buildtools
    git
    go
    jq
    pre-commit
  ];

  # 2. Define the Profiles
  profiles = {
    # The 'ci' profile: only adds what's strictly necessary for the runner
    ci.module = {
      packages = with pkgs; [
        lcov
      ];
      # You can even disable expensive checks in CI if needed
      git-hooks.hooks.shellcheck.enable = false;
    };

    # The 'dev' profile: adds interactive tools and quality-of-life packages
    dev.module = {
      packages = with pkgs; [
        bat
        difftastic
        dua
        eza
        golangci-lint
        just
        kondo
        lazygit
        lcov
        mdcat
        openssl
        python3
        readline
        ripgrep
        starship
        tokei
        uv
        vim
        zlib
      ];

      enterShell = ''
        alias ls='eza --icons'
        eval "$(starship init bash)"
      '';
    };
  };

  # Global environment variables
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
}
