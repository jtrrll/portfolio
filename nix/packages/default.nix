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
                src = pkgs.nix-gitignore.gitignoreRecursiveSource [ ] ../../go;
                vendorHash = "sha256-U53wKtH8I9ESFb6QiTvOi4Ha8R216EZjX+3EuiWjq5I=";
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
