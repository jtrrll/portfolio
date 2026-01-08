{
  description = "Jackson Terrill's personal portfolio.";

  inputs = {
    ### Development dependencies ###
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    devenv = {
      inputs.nixpkgs.follows = "nixpkgs";
      url = "github:cachix/devenv";
    };
    flake-parts = {
      inputs.nixpkgs-lib.follows = "nixpkgs";
      url = "github:hercules-ci/flake-parts";
    };
    justix = {
      inputs.nixpkgs.follows = "nixpkgs";
      url = "github:jtrrll/justix";
    };
    snekcheck = {
      inputs.nixpkgs.follows = "nixpkgs";
      url = "github:jtrrll/snekcheck";
    };
    treefmt-nix = {
      inputs.nixpkgs.follows = "nixpkgs";
      url = "github:numtide/treefmt-nix";
    };

    ### Build dependencies ###
    nixpkgs-resume.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    nixpkgs-server.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs =
    { flake-parts, ... }@inputs:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        ./apps
        ./checks
        ./dev_shells
        ./formatter
        ./packages
      ];
      systems = inputs.nixpkgs.lib.systems.flakeExposed;
    };
}
