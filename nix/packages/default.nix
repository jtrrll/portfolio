{
  perSystem =
    { pkgs, ... }:
    {
      packages = {
        resume = pkgs.callPackage ./resume.nix { };
      };
    };
}
