{
  description = "A flake to build Snake Go for different targets";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  # I want to create multiple targets for some distros and windows. Distros will make a packages
  outputs = { self, nixpkgs }:
    flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs { inherit system; };
    in with pkgs; rec {

    });
}
