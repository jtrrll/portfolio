{ inputs, ... }:
{
  perSystem =
    {
      system,
      ...
    }:
    let
      resumePkgs = inputs.nixpkgs-resume.legacyPackages.${system};
    in
    {
      packages.resume = resumePkgs.callPackage (
        {
          ibm-plex,
          typst,
          stdenvNoCC,
        }:
        let
          typstWithPackages = typst.withPackages (typstPkgs: with typstPkgs; [ basic-resume_0_2_9 ]);
        in
        stdenvNoCC.mkDerivation {
          name = "jackson-terrill-resume";
          src = ./resume.typ;
          nativeBuildInputs = [
            ibm-plex
            typstWithPackages
          ];
          passthru = { inherit typst typstWithPackages; };

          dontUnpack = true;
          buildPhase = ''
            typst compile --font-path ${ibm-plex} $src resume.pdf
          '';
          installPhase = ''
            cp resume.pdf $out
          '';
        }
      ) { };
    };
}
