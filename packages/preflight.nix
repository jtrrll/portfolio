{
  fetchFromGitHub,
  runCommand,
}:
let
  tailwindcss = fetchFromGitHub {
    owner = "tailwindlabs";
    repo = "tailwindcss";
    rev = "v4.1.14";
    hash = "sha256-BGySdbLTvZ40i4LMkyXv+aD79p050tD2r/s1G3tGMfc=";
  };
in
runCommand "preflight.css" { } ''
  cp ${tailwindcss}/packages/tailwindcss/preflight.css $out
''
