{ inputs, ... }:
{
  imports = [ inputs.treefmt-nix.flakeModule ];
  perSystem = _: {
    treefmt.programs = {
      actionlint.enable = true;
      deadnix.enable = true;
      gofmt.enable = true;
      keep-sorted.enable = true;
      nixfmt.enable = true;
      statix.enable = true;
      templ.enable = true;
      typstyle.enable = true;
      yamlfmt.enable = true;
    };
  };
}
