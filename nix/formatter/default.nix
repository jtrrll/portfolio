{ inputs, ... }:
{
  imports = [ inputs.treefmt-nix.flakeModule ];
  perSystem = _: {
    treefmt.programs = builtins.addErrorContext "while defining formatter" {
      actionlint.enable = true;
      deadnix.enable = true;
      gofmt.enable = true;
      nixfmt.enable = true;
      statix.enable = true;
      templ.enable = true;
      typstyle.enable = true;
      yamlfmt.enable = true;
    };
  };
}
