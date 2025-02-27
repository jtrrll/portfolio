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
      nativeBuildInputs = [
        inputs.templ.packages.${system}.templ
        pkgs.makeWrapper
      ];
      pname = name;
      postInstall = ''
        wrapProgram $out/bin/server \
        --set OTEL_RESOURCE_ATTRIBUTES "service.name=${name},service.version=${version}"
      '';
      preBuild = ''
        templ generate -log-level error
      '';
      src = ../go;
    };
  in {
    apps.server.program = "${pkg}/bin/server";
    packages.${name} = pkg;
  };
}
