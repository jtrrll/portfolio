{ inputs, ... }:
{
  imports = [ inputs.devenv.flakeModule ];
  perSystem =
    {
      inputs',
      lib,
      pkgs,
      self',
      ...
    }:
    {
      devenv = {
        modules = [
          inputs.justix.devenvModules.default
          {
            containers = lib.mkForce { }; # Workaround to remove containers from flake checks.
          }
        ];
        shells.default =
          { config, ... }:
          {
            enterShell = lib.getExe (
              pkgs.writeShellApplication rec {
                meta.mainProgram = name;
                name = "splashScreen";
                runtimeInputs = [
                  pkgs.lolcat
                  pkgs.uutils-coreutils-noprefix
                ];
                text = ''
                  printf "
                   ▐▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄  ▄▄▌  ▄▄▌
                    ·██•██  ▀▄ █·▀▄ █·██•  ██•
                  ▪▄ ██ ▐█.▪▐▀▀▄ ▐▀▀▄ ██▪  ██▪
                  ▐▌▐█▌ ▐█▌·▐█•█▌▐█•█▌▐█▌▐▌▐█▌▐▌
                   ▀▀▀• ▀▀▀ .▀  ▀.▀  ▀.▀▀▀ .▀▀▀\n" | lolcat
                  printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n"
                '';
              }
            );

            claude.code.enable = true;

            justix = {
              enable = true;
              mcpServer.enable = true;
              justfile.config.recipes = {
                build-assets = {
                  attributes = {
                    doc = "Builds static assets for the server";
                    private = true;
                  };
                  commands =
                    let
                      assetDir = "${config.devenv.root}/go/cmd/server/static";
                    in
                    ''
                      @cp --force $(nix build --no-link --print-out-paths .#preflight) ${assetDir}/preflight.css
                      @cp --force $(nix build --no-link --print-out-paths .#resume) ${assetDir}/jackson_terrill_resume.pdf
                    '';
                };
                default = {
                  attributes = {
                    default = true;
                    doc = "Lists available recipes";
                    private = true;
                  };
                  commands = "@just --list";
                };
                develop-server = {
                  aliases = [ "dev" ];
                  attributes = {
                    doc = "Runs the server in development mode with live reloading";
                    working-directory = "go";
                  };
                  commands = ''
                    @templ generate --watch --proxy="http://localhost:8080" --cmd="go run -tags dev ./cmd/server"
                  '';
                  dependencies = [ "build-assets" ];
                };
                fmt = {
                  attributes.doc = "Formats and lints files";
                  commands = ''
                    @find "{{ paths }}" ! -path '*/.*' -exec ${lib.getExe inputs'.snekcheck.packages.default} --fix {} +
                    @nix fmt -- {{ paths }}
                  '';
                  parameters = [ "*paths='.'" ];
                };
              };
            };

            languages = {
              go = {
                enable = true;
                package = self'.packages.default.go;
              };
              nix.enable = true;
              typst = {
                enable = true;
                package = self'.packages.resume.typst;
              };
            };

            packages = [
              self'.packages.default.templ
              pkgs.woff2
            ];

            services.opentelemetry-collector.enable = true;

            git-hooks = {
              default_stages = [ "pre-push" ];
              hooks = {
                actionlint.enable = true;
                check-added-large-files = {
                  enable = true;
                  stages = [ "pre-commit" ];
                };
                check-json.enable = true;
                check-yaml.enable = true;
                deadnix.enable = true;
                detect-private-keys = {
                  enable = true;
                  stages = [ "pre-commit" ];
                };
                end-of-file-fixer.enable = true;
                flake-checker.enable = true;
                markdownlint.enable = true;
                mixed-line-endings.enable = true;
                nil.enable = true;
                no-commit-to-branch = {
                  enable = true;
                  stages = [ "pre-commit" ];
                };
                ripsecrets = {
                  enable = true;
                  stages = [ "pre-commit" ];
                };
                statix.enable = true;
              };
            };
          };
      };
    };
}
