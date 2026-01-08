{ inputs, ... }:
{
  perSystem =
    {
      system,
      ...
    }:
    let
      resumePkgs = inputs.nixpkgs-resume.legacyPackages.${system};
      serverPkgs = inputs.nixpkgs-server.legacyPackages.${system};
    in
    {
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
                vendorHash = "sha256-G4W95LmSJaMEe9gJr4jYEgTKnRylm1SnPNPr/+xEfY0=";
              } args
            );
        in
        rec {
          default = serverPkgs.callPackage ./server.nix {
            inherit buildGoModule preflight resume;
          };
          preflight = serverPkgs.callPackage ./preflight.nix { };
          resume = resumePkgs.callPackage ./resume { };
        };
    };
}
