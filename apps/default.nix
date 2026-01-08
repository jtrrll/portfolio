{
  perSystem =
    { lib, self', ... }:
    {
      apps.default = {
        type = "app";
        program = lib.getExe self'.packages.default;
      };
    };
}
