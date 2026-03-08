{
  description = "Development environment using devenv, equivalent to the original shell.nix";

  inputs = {
    nixpkgs.url = "https://github.com/NixOS/nixpkgs/archive/nixos-25.11.tar.gz";
    devenv.url = "github:cachix/devenv";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      nixpkgs,
      devenv,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = devenv.lib.mkShell {
          inherit pkgs;
          modules = [ ./devenv.nix ];
        };
      }
    );
}
