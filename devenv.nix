{
  pkgs,
  ...
}:

{
  # Packages used everywhere (Dev + CI)
  packages = with pkgs; [
    bazelisk
    bazel-buildtools
  ];

  scripts.bazel.exec = "bazelisk \"$@\"";

  # languages.rust = {
  #   enable = true;
  # };

  # 2. Define the Profiles
  profiles = {
    # The 'ci' profile: only adds what's strictly necessary for the runner
    ci.module = {
      packages = with pkgs; [
        lcov
      ];

      enterShell = ''
        alias bazel='bazelisk'
      '';

      # You can even disable expensive checks in CI if needed
      git-hooks.hooks.shellcheck.enable = false;
    };

    # The 'dev' profile: adds interactive tools and quality-of-life packages
    dev.module = {
      packages = with pkgs; [
        bat
        blesh
        difftastic
        dua
        eza
        git
        google-java-format
        jq
        just
        kondo
        lazygit
        lcov
        mdcat
        onefetch
        openssl
        perl540Packages.PerlTidy
        podman
        python3
        readline
        ripgrep
        rustlings
        starship
        tokei
        uv
        vim
        workshop-runner
        zlib
      ];

      # difftastic.enable = true;

      enterShell = ''
        alias ls='eza --icons'

        source "${pkgs.blesh}/share/blesh/ble.sh"

        if [[ -n "$GHCR_PAT" ]]; then
          echo "Logging into GitHub Container Registry with provided GHCR_PAT"
          echo $GHCR_PAT | podman login ghcr.io -u racosta --password-stdin
        fi

        eval "$(starship init bash)"

        onefetch --nerd-fonts --number-of-languages=8
      '';

      git-hooks = {
        package = pkgs.prek;
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
          clang-format = {
            enable = true;
            types_or = [
              "c"
              "c++"
            ];
          };
          deadnix.enable = true;
          denofmt.enable = true;
          denolint.enable = true;
          end-of-file-fixer.enable = true;
          # eslint.enable = true;
          flake8 = {
            enable = true;
            args = [
              "--max-line-length=120"
            ];
          };
          gofmt.enable = true;
          golangci-lint.enable = true;
          golines.enable = true;
          google-java-format = {
            enable = true;
            description = "Format Java code with google-java-format";
            entry = "${pkgs.google-java-format}/bin/google-java-format -i --set-exit-if-changed";
            pass_filenames = true;
            types = [ "java" ];
          };
          govet.enable = true;
          isort = {
            enable = true;
            args = [
              "--profile=black"
              "--filter-files"
            ];
          };
          keep-sorted = {
            enable = true;
            description = "Ensure sections of files are sorted using keep-sorted";
            entry = "${pkgs.keep-sorted}/bin/keep-sorted";
            pass_filenames = true;
            files = "BUILD|WORKSPACE|\\.(bazel|bzl|nix|in|txt|md|rs|py|go)$";
          };
          markdownlint.enable = true;
          # projects/python_calculator/calculator_test.py:3: error: Cannot find implementation or library stub for module named "calculator"  [import-not-found]
          # projects/python_calculator/calculator_test.py:3: note: See https://mypy.readthedocs.io/en/stable/running_mypy.html#missing-imports
          # Found 1 error in 1 file (checked 4 source files)
          # projects/python_web/main.py:3: error: Cannot find implementation or library stub for module named "flask"  [import-not-found]
          # projects/python_web/main.py:3: note: See https://mypy.readthedocs.io/en/stable/running_mypy.html#missing-imports
          # Found 1 error in 1 file (checked 2 source files)
          mypy.enable = false;
          nixfmt.enable = true;
          no-commit-to-branch = {
            enable = true;
            description = "Protect branches from direct commits";
            entry =
              let
                script = pkgs.writeShellScript "precommit-no-commit-to-branch" ''
                  set -e
                  current_branch=$(${pkgs.git}/bin/git symbolic-ref --short HEAD)

                  # Check if the current branch is 'main'
                  if [[ "$current_branch" = "main" ]]; then
                    echo "Error: Cannot commit directly to the 'main' branch."
                    echo "Please switch to a feature branch or use a pull request."
                    exit 1 # Exit with a non-zero status to abort the commit
                  fi
                '';
              in
              builtins.toString script;
            pass_filenames = false;
          };
          perlcritic = {
            enable = true;
            description = "Run Perl::Critic linter on Perl code";
            entry = "${pkgs.perl540Packages.PerlCritic}/bin/perlcritic";
            args = [
              "--profile"
              "/dev/null"
              "--severity"
              "5"
            ];
            types = [ "perl" ];
          };
          perltidy = {
            enable = true;
            description = "Format Perl code with perltidy";
            entry = "${pkgs.perl540Packages.PerlTidy}/bin/perltidy";
            args = [
              "--noprofile"
              "--perl-best-practices"
              "--nostandard-output"
              "--warning-output"
              "--backup-and-modify-in-place"
              "--backup-file-extension=/"
              "--indent-columns=2"
              "--maximum-line-length=120"
            ];
            types = [ "perl" ];
          };
          pydocstyle = {
            enable = true;
            description = "Run pydocstyle linter on Python code";
            entry = "${pkgs.python313Packages.pydocstyle}/bin/pydocstyle --convention=google";
            types = [ "python" ];
          };
          revive.enable = true;
          ruff.enable = true;
          ruff-format.enable = true;
          rustfmt = {
            enable = true;
            description = "Format Rust code with rustfmt";
            entry = "${pkgs.bazel_8}/bin/bazel run @rules_rust//:rustfmt";
            types = [ "rust" ];
          };
          shellcheck.enable = true;
          shfmt = {
            enable = true;
            description = "Format shell files.";
            types = [ "shell" ];
            entry = "${pkgs.shfmt}/bin/shfmt -w -i 2 -l -s";
          };
          staticcheck.enable = true;
          trim-trailing-whitespace.enable = true;
          # ty = {
          #   enable = true;
          #   description = "Run 'ty check' for extremely fast Python type checking.";
          #   entry = "${pkgs.ty}/bin/ty check";
          #   types = [ "python" ];
          #   require_serial = true;
          # };
          update-cargo-lock = {
            enable = true;
            description = "Update Cargo lock file";
            entry = "${pkgs.cargo}/bin/cargo generate-lockfile";
            files = "^third_party/rust/Cargo\\.toml$";
            pass_filenames = false;
          };
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
  env.STARSHIP_LOG = "ERROR";
}
