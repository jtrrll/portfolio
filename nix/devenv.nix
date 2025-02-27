{inputs, ...}: {
  imports = [
    inputs.devenv.flakeModule
  ];
  perSystem = {
    config,
    lib,
    pkgs,
    system,
    ...
  }: {
    devenv = {
      modules = [
        inputs.env-help.devenvModule
      ];
      shells.default = let
        buildInputs = config.packages.portfolio.nativeBuildInputs;
        goPkg = lib.findFirst (pkg: builtins.match "go" pkg.pname != null) pkgs.go buildInputs;
      in {
        enterShell = "${pkgs.writeShellApplication {
          name = "splashScreen";
          runtimeInputs = [
            pkgs.lolcat
            pkgs.uutils-coreutils-noprefix
          ];
          text = ''
            printf " ▐▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄  ▄▄▌  ▄▄▌
              ·██•██  ▀▄ █·▀▄ █·██•  ██•
            ▪▄ ██ ▐█.▪▐▀▀▄ ▐▀▀▄ ██▪  ██▪
            ▐▌▐█▌ ▐█▌·▐█•█▌▐█•█▌▐█▌▐▌▐█▌▐▌
             ▀▀▀• ▀▀▀ .▀  ▀.▀  ▀.▀▀▀ .▀▀▀ \n" | lolcat
            printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n"
          '';
        }}/bin/splashScreen";

        env = let
          PROJECT_ROOT = config.devenv.shells.default.env.DEVENV_ROOT;
        in {
          inherit PROJECT_ROOT;
          GO_ROOT = "${PROJECT_ROOT}/go";
          NIX_ROOT = "${PROJECT_ROOT}/nix";
        };

        env-help.enable = true;

        languages = {
          go = {
            enable = true;
            package = goPkg;
          };
          nix.enable = true;
        };

        packages = [
          inputs.templ.packages.${system}.templ
        ];

        pre-commit = {
          default_stages = ["pre-push"];
          hooks = {
            actionlint.enable = true;
            check-added-large-files = {
              enable = true;
              stages = ["pre-commit"];
            };
            check-yaml.enable = true;
            deadnix.enable = true;
            detect-private-keys = {
              enable = true;
              stages = ["pre-commit"];
            };
            end-of-file-fixer.enable = true;
            flake-checker.enable = true;
            lint = {
              enable = true;
              entry = "lint";
              name = "lint";
              pass_filenames = false;
            };
            markdownlint.enable = true;
            mixed-line-endings.enable = true;
            nil.enable = true;
            no-commit-to-branch = {
              enable = true;
              stages = ["pre-commit"];
            };
            ripsecrets = {
              enable = true;
              stages = ["pre-commit"];
            };
            statix.enable = true;
          };
        };

        scripts = {
          bench = {
            description = "Runs all benchmark tests.";
            exec = "${pkgs.writeShellApplication {
              name = "build";
              runtimeInputs = [
                goPkg
                inputs.templ.packages.${system}.templ
              ];
              text = ''
                cd "$GO_ROOT" && \
                templ generate -log-level error
                go test ./... -bench=.
              '';
            }}/bin/bench";
          };
          build = {
            description = "Builds the project's binaries.";
            exec = "${pkgs.writeShellApplication {
              name = "build";
              runtimeInputs = [
                goPkg
                inputs.gomod2nix.legacyPackages.${system}.gomod2nix
              ];
              text = ''
                (cd "$GO_ROOT" && go mod tidy && gomod2nix) && \
                nix build "$PROJECT_ROOT"#portfolio
              '';
            }}/bin/build";
          };
          e2e = {
            description = "Runs all end-to-end tests.";
            exec = "${pkgs.writeShellApplication {
              name = "e2e";
              text = ''
                build && \
                echo "TODO"
              '';
            }}/bin/e2e";
          };
          lint = {
            description = "Lints the project.";
            exec = "${pkgs.writeShellApplication {
              name = "lint";
              runtimeInputs = [
                goPkg
                inputs.snekcheck.packages.${system}.snekcheck
                inputs.templ.packages.${system}.templ
                pkgs.golangci-lint
              ];
              text = ''
                snekcheck --fix "$PROJECT_ROOT" && \
                nix fmt "$PROJECT_ROOT/flake.nix" "$NIX_ROOT" -- --quiet && \
                (cd "$GO_ROOT" && \
                templ generate -log-level error
                go mod tidy && go fmt ./... && go vet ./... && \
                golangci-lint run ./...)
              '';
            }}/bin/lint";
          };
          run-server = {
            description = "Runs the portfolio server.";
            exec = "${pkgs.writeShellApplication {
              name = "run-server";
              runtimeInputs = [
                goPkg
                inputs.gomod2nix.legacyPackages.${system}.gomod2nix
              ];
              text = ''
                (cd "$GO_ROOT" && go mod tidy && gomod2nix) && \
                nix run "$PROJECT_ROOT"#server -- "$@"
              '';
            }}/bin/run-server";
          };
          unit = {
            description = "Runs all unit tests.";
            exec = "${pkgs.writeShellApplication {
              name = "unit";
              runtimeInputs = [
                goPkg
                inputs.templ.packages.${system}.templ
              ];
              text = ''
                cd "$GO_ROOT" && \
                templ generate -log-level error && \
                go test --cover ./...
              '';
            }}/bin/unit";
          };
        };
      };
    };
  };
}
