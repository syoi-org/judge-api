{
  description = "SYOI Online Judge";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix.url = "github:nix-community/gomod2nix";
    gomod2nix.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils, gomod2nix }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = gomod2nix.legacyPackages.${system}.buildGoApplication {
          pname = "judge-api";
          name = "judge-api";
          src = ./.;
          modules = ./gomod2nix.toml;
        };
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            air
            go
            cobra-cli
            just
          ] ++ [
            gomod2nix.packages.${system}.default
          ];
        };
      });
}
