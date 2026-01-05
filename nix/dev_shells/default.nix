{ inputs, ... }:
{
  imports = [ inputs.devenv.flakeModule ];
  perSystem =
    {
      inputs',
      lib,
      pkgs,
      self',
      system,
      ...
    }:
    {
      devenv = {
        modules = [
          inputs.justix.devenvModules.default
          {
            containers = lib.mkForce { }; # Workaround to remove containers from flake checks.
          }
          {
            claude.code.enable = true;
            justix = {
              enable = true;
              mcpServer.enable = true;
              justfile.config.recipes = {
                default = {
                  attributes = {
                    default = true;
                    doc = "Lists available recipes";
                    private = true;
                  };
                  commands = "@just --list";
                };
                develop-server = {
                  attributes.doc = "Runs the server in development mode with live reloading";
                  commands = ''
                    @templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
                  '';
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

            languages = {
              go = {
                enable = true;
                package = self'.packages.default.go;
              };
              nix.enable = true;
              typst.enable = true;
            };

            packages = [
              self'.packages.default.templ
              pkgs.woff2
            ];

            scripts.dev-server = {
              exec = ''
                set -euo pipefail

                find ${config.devenv.root}/go | entr -n sh -c '
                  log() {
                    echo "[dev-server][$(date +'%H:%M:%S')] $*"
                  }

                  log "Building new version..."
                  nix build --impure --expr "(builtins.getFlake (toString ./.)).packages.${system}.default.override { dev = true; }" -o /tmp/portfolio-server-build --quiet &> /dev/null || {
                    log "❌ Build failed, keeping current server running"
                    exit 0
                  }

                  if [ -f /tmp/portfolio-server.pid ]; then
                    PID=$(cat /tmp/portfolio-server.pid)
                    if kill -0 "$PID" 2>/dev/null; then
                      log "Stopping old server (PID $PID)..."
                      kill "$PID"
                      wait "$PID" 2>/dev/null || true
                    fi
                  fi

                  log "Starting new server..."
                  /tmp/portfolio-server-build/bin/server run --port 8080 &
                  echo $! > /tmp/portfolio-server.pid
                  log "✅ Running on http://localhost:8080"
                '
              '';
              packages = [ pkgs.entr ];
              description = "Runs the server and rebuilds on file changes";
            };

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
