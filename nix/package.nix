{inputs, ...}: {
  perSystem = {
    pkgs,
    system,
    ...
  }: let
    name = "portfolio";
    version = "0.0";
    pkg = inputs.gomod2nix.legacyPackages.${system}.buildGoApplication {
      inherit version;
      modules = ../go/gomod2nix.toml;
      nativeBuildInputs = [pkgs.makeWrapper];
      pname = name;
      postInstall = ''
        wrapProgram $out/bin/server \
        --set OTEL_RESOURCE_ATTRIBUTES "service.name=${name},service.version=${version}"
      '';
      preBuild = ''
        ${inputs.templ.packages.${system}.templ}/bin/templ generate
      '';
      src = ../go;
    };
  in {
    apps.server.program = "${pkg}/bin/server";
    packages.${name} = pkg;
  };
}
