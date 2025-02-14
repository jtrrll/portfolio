{inputs, ...}: {
  imports = [
    ./devenv.nix
    ./package.nix
  ];
  perSystem = {pkgs, ...}: {
    formatter = pkgs.alejandra;
  };
  systems = inputs.nixpkgs.lib.systems.flakeExposed;
}
