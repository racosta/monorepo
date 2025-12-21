{
  pkgs,
  ...
}:

{
  # Packages used everywhere (Dev + CI)
  packages = with pkgs; [
    bazel_8
    bazel-buildtools
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
        blesh
        buildifier
        difftastic
        dua
        eza
        git
        go
        gofumpt
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

      enterShell = ''
        alias ls='eza --icons'

        source "${pkgs.blesh}/share/blesh/ble.sh"

        eval "$(starship init bash)"
      '';

      git-hooks.package = pkgs.prek;
      git-hooks = {
        enable = true;

        hooks = {
          action-validator.enable = true;
          black.enable = true;
          buildifier = {
            enable = true;
            package = pkgs.buildifier;
            description = "Format Bazel files with buildifier";
            entry = "${pkgs.buildifier}/bin/buildifier";
            types = [ "bazel" ];
          };
          check-added-large-files.enable = true;
          check-json.enable = true;
          check-merge-conflicts.enable = true;
          check-symlinks.enable = true;
          check-toml.enable = true;
          check-yaml.enable = true;
          # The checkmake tool does not support include directives well,
          # so creates many false positives
          checkmake.enable = false;
          deadnix.enable = true;
          end-of-file-fixer.enable = true;
          flake8 = {
            enable = true;
            args = [
              "--max-line-length=120"
            ];
          };
          gofmt.enable = true;
          golangci-lint.enable = true;
          golines.enable = true;
          govet.enable = true;
          isort = {
            enable = true;
            args = [
              "--profile=black"
              "--filter-files"
            ];
          };
          markdownlint.enable = true;
          # projects/python_calculator/calculator_test.py:3: error: Cannot find implementation or library stub for module named "calculator"  [import-not-found]
          # projects/python_calculator/calculator_test.py:3: note: See https://mypy.readthedocs.io/en/stable/running_mypy.html#missing-imports
          # Found 1 error in 1 file (checked 4 source files)
          # projects/python_web/main.py:3: error: Cannot find implementation or library stub for module named "flask"  [import-not-found]
          # projects/python_web/main.py:3: note: See https://mypy.readthedocs.io/en/stable/running_mypy.html#missing-imports
          # Found 1 error in 1 file (checked 2 source files)
          #mypy.enable = true;
          nixfmt.enable = true;
          revive.enable = true;
          shellcheck.enable = true;
          shfmt.enable = true;
          staticcheck.enable = true;
          trim-trailing-whitespace.enable = true;
          yamlfmt = {
            enable = true;
            settings = {
              configPath = ".yamlfmt.yaml";
              lint-only = false;
            };
          };
          yamllint = {
            enable = true;
            args = [
              "--format=parsable"
              "--strict"
            ];
            settings.configPath = ".yamllint.yaml";
          };
        };
      };
    };
  };

  # Global environment variables
  env.GREETING = "Hello, Nix!";
}
