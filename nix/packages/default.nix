{
  perSystem =
    { pkgs, ... }:
    {
      packages = rec {
        default = pkgs.callPackage ./server.nix {
          inherit resume;
        };
        resume = pkgs.callPackage ./resume.nix { };
      };
    };
}
