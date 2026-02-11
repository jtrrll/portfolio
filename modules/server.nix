{ inputs, ... }:
{
  perSystem =
    {
      lib,
      self',
      system,
      ...
    }:
    let
      serverPkgs = inputs.nixpkgs-server.legacyPackages.${system};
    in
    {
      apps.default = {
        type = "app";
        program = lib.getExe self'.packages.default;
      };
      packages =
        let
          buildGoModule =
            args:
            serverPkgs.buildGoModule (
              serverPkgs.lib.recursiveUpdate {
                meta = {
                  homepage = "https://github.com/jtrrll/portfolio";
                  license = serverPkgs.lib.licenses.mit;
                };
                src = serverPkgs.lib.cleanSource (serverPkgs.nix-gitignore.gitignoreRecursiveSource [ ] ../go);
                vendorHash = "sha256-Iz2kqWXZv+0M5BAh400ZwxQe6wYkmwtKQj26fiTK2P0=";
              } args
            );
        in
        {
          default =
            serverPkgs.callPackage
              (
                {
                  buildGoModule,
                  preflight,
                  resume,
                  templ,
                }:
                buildGoModule {
                  pname = "portfolio-server";
                  version = "0.0.0";

                  meta = {
                    description = "Jackson Terrill's personal portfolio.";
                    mainProgram = "server";
                  };
                  subPackages = [ "cmd/server" ];
                  nativeBuildInputs = [ templ ];
                  passthru = {
                    inherit templ;
                  };

                  preBuild = ''
                    templ generate
                    cp ${preflight} cmd/server/static/preflight.css
                    cp ${resume} cmd/server/static/jackson_terrill_resume.pdf
                  '';

                  env.CGO_ENABLED = 0;
                }
              )
              {
                inherit buildGoModule;
                inherit (self'.packages) preflight resume;
              };
        };
    };
}
