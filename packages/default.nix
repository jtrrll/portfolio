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
                vendorHash = "sha256-s+zOIpUQ9t19LJDiOhZou3o1e3RWdaHZZn2FawBsAJY=";
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
