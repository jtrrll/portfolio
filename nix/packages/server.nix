{
  buildGoModule,
  dev ? false,
  fetchFromGitHub,
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
    mainProgram = "server";
  };
  subPackages = [ "cmd/server" ];
  tags = if dev then [ "dev" ] else [ ];
  nativeBuildInputs = [ templ ];
  passthru = {
    inherit templ;
  };

  preBuild = ''
    templ generate
    cp ${preflight} cmd/server/static/preflight.css
    cp ${resume} cmd/server/static/jackson_terrill_resume.pdf
  '';
}
