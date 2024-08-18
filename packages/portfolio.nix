{pkgs, ...}: {
  packages."portfolio" = let
    cargoTOML = builtins.fromTOML (builtins.readFile ../Cargo.toml);
  in
    pkgs.rustPlatform.buildRustPackage {
      inherit (cargoTOML.package) version;
      cargoLock.lockFile = ../Cargo.lock;
      meta = {
        description = "jtrrll's personal portfolio";
        homepage = "https://github.com/jtrrll/portfolio";
        license = pkgs.lib.licenses.mit;
        mainProgram = "portfolio";
      };
      pname = cargoTOML.package.name;
      src = ./..;
    };
}
