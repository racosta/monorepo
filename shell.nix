{
  ch2505 ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-25.05.tar.gz") {}
}:

ch2505.mkShell {
  packages = with ch2505; [
    bat
    bazel_7
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
