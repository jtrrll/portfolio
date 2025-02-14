{inputs, ...}: {
  perSystem = {system, ...}: let
    pkg = inputs.gomod2nix.legacyPackages.${system}.buildGoApplication {
      modules = ../go/gomod2nix.toml;
      pname = "portfolio";
      src = ../go;
      version = "0.0";
    };
  in {
    apps.server.program = "${pkg}/bin/server";
    packages.portfolio = pkg;
  };
}
