{
  ch2511 ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-25.11.tar.gz") {}
}:

ch2511.mkShell {
  packages = with ch2511; [
    bat
    bazel_8
    bazel-buildtools
    difftastic
    #direnv
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

  GREETING = "Hello, Nix!";

  shellHook = ''
    alias ls='eza --icons'

    eval "$(starship init bash)"
  '';
}
