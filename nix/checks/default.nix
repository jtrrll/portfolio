{
  perSystem =
    {
      inputs',
      pkgs,
      ...
    }:
    {
      checks = builtins.addErrorContext "while defining checks" {
        snekcheck =
          pkgs.runCommandLocal "snekcheck"
            {
              buildInputs = [ inputs'.snekcheck.packages.default ];
            }
            ''
              find ../../** -exec snekcheck {} +
              touch $out
            '';
      };
    };
}
