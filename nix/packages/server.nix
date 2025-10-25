{
  buildGoModule,
  fetchFromGitHub,
  lib,
  resume,
  templ,
}:
let
  preflight = "${
    (fetchFromGitHub {
      owner = "tailwindlabs";
      repo = "tailwindcss";
      rev = "v4.1.14";
      hash = "sha256-BGySdbLTvZ40i4LMkyXv+aD79p050tD2r/s1G3tGMfc=";
    })
  }/packages/tailwindcss/preflight.css";
in
buildGoModule {
  pname = "portfolio-server";
  version = "0.0.0";

  meta = {
    description = "Jackson Terrill's personal portfolio.";
    homepage = "https://github.com/jtrrll/portfolio";
    license = lib.licenses.mit;
    mainProgram = "server";
  };
  src = lib.cleanSourceWith {
    filter = absPath: _: !(lib.strings.hasSuffix "_templ.go" absPath);
    src = lib.cleanSource ../../go;
  };
  subPackages = [ "cmd/server" ];
  nativeBuildInputs = [ templ ];
  passthru = {
    inherit templ;
  };
  vendorHash = "sha256-U53wKtH8I9ESFb6QiTvOi4Ha8R216EZjX+3EuiWjq5I=";

  preBuild = ''
    templ generate
    cp ${preflight} cmd/server/static/preflight.css
    cp ${resume} cmd/server/static/jackson_terrill_resume.pdf
  '';
}
