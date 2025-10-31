{
  perSystem =
    { lib, pkgs, ... }:
    {
      packages =
        let
          buildGoModule =
            args:
            pkgs.buildGoModule (
              lib.recursiveUpdate {
                meta = {
                  homepage = "https://github.com/jtrrll/portfolio";
                  license = lib.licenses.mit;
                };
                src = lib.cleanSource (pkgs.nix-gitignore.gitignoreRecursiveSource [ ] ../../go);
                vendorHash = "sha256-G4W95LmSJaMEe9gJr4jYEgTKnRylm1SnPNPr/+xEfY0=";
              } args
            );
        in
        rec {
          default = pkgs.callPackage ./server.nix {
            inherit buildGoModule resume;
          };
          resume = pkgs.callPackage ./resume.nix { };
        };
    };
}
