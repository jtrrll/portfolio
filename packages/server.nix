{
  buildGoModule,
  preflight,
  resume,
  templ,
}:
buildGoModule {
  pname = "portfolio-server";
  version = "0.0.0";

  meta = {
    description = "Jackson Terrill's personal portfolio.";
    mainProgram = "server";
  };
  subPackages = [ "cmd/server" ];
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
