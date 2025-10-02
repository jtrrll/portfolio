{
  ibm-plex,
  typst,
  stdenvNoCC,
}:
let
  typstWithPackages = typst.withPackages (typstPkgs: with typstPkgs; [ basic-resume_0_2_7 ]);
in
stdenvNoCC.mkDerivation {
  name = "jackson-terrill-resume";
  src = ../../typst/resume.typ;
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
