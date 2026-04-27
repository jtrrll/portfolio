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
        meta = {
          inherit (self'.packages.default.meta) description;
        };
        type = "app";
        program = lib.getExe self'.packages.default;
      };
      packages.default =
        serverPkgs.callPackage
          (
            {
              preflight,
              resume,
              templ,
              testers,
              versionCheckHook,
            }:
            serverPkgs.buildGoModule (finalAttrs: {
              pname = "portfolio";
              version = "0.0.0";
              src = serverPkgs.lib.cleanSource (serverPkgs.nix-gitignore.gitignoreRecursiveSource [ ] ../go);
              vendorHash = "sha256-UABDY1dtEel9/eUBasEcZQIWB3FBMWXDA7znbPQY8as=";

              meta = {
                description = "Jackson Terrill's personal portfolio.";
                homepage = "https://github.com/jtrrll/portfolio";
                license = serverPkgs.lib.licenses.mit;
                mainProgram = "server";
                platforms = lib.platforms.all;
              };
              passthru = {
                inherit templ;
                tests.version = testers.testVersion { package = finalAttrs.finalPackage; };
              };

              env.CGO_ENABLED = 0;

              nativeBuildInputs = [ templ ];
              preBuild = ''
                templ generate
                cp ${preflight} cmd/server/static/preflight.css
                cp ${resume} cmd/server/static/jackson_terrill_resume.pdf
              '';
              doInstallCheck = true;
              nativeInstallCheckInputs = [ versionCheckHook ];
            })
          )
          {
            inherit (self'.packages) preflight resume;
          };
    };
}
