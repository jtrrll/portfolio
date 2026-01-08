{ self, ... }:
{
  perSystem =
    {
      inputs',
      pkgs,
      ...
    }:
    {
      checks = {
        snekcheck =
          pkgs.runCommandLocal "snekcheck"
            {
              buildInputs = [ inputs'.snekcheck.packages.default ];
            }
            ''
              find ${self}/** -exec snekcheck {} +
              touch $out
            '';
      };
    };
}
