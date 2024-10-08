{inputs, ...}: {
  imports = [
    inputs.devenv.flakeModule
  ];
  perSystem = {pkgs, ...}: {
    devenv = {
      modules = [
        inputs.env-help.devenvModule
      ];
      shells.default = {
        enterShell = ''
          printf " ▄▄▄·      ▄▄▄  ▄▄▄▄▄·▄▄▄      ▄▄▌  ▪
          ▐█ ▄█▪     ▀▄ █·•██  ▐▄▄·▪     ██•  ██ ▪
          ██▀· ▄█▀▄ ▐▀▀▄  ▐█.▪██▪  ▄█▀▄ ██▪  ▐█· ▄█▀▄
          ▐█▪·•▐█▌.▐▌▐█•█▌ ▐█▌·██▌.▐█▌.▐▌▐█▌▐▌▐█▌▐█▌.▐▌
          .▀    ▀█▄▀▪.▀  ▀ ▀▀▀ ▀▀▀  ▀█▄▀▪.▀▀▀ ▀▀▀ ▀█▄▀▪\n" | ${pkgs.lolcat}/bin/lolcat
          printf "\033[0;1;36mDEVSHELL ACTIVATED\033[0m\n"
        '';

        env-help.enable = true;

        languages = {
          rust.enable = true;
          nix.enable = true;
        };

        packages = [
          pkgs.commitizen
        ];

        pre-commit = {
          default_stages = ["pre-push"];
          hooks = {
            actionlint.enable = true;
            alejandra.enable = true;
            check-added-large-files = {
              enable = true;
              stages = ["pre-commit"];
            };
            check-yaml.enable = true;
            clippy.enable = true;
            commitizen.enable = true;
            deadnix.enable = true;
            detect-private-keys = {
              enable = true;
              stages = ["pre-commit"];
            };
            end-of-file-fixer.enable = true;
            flake-checker.enable = true;
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
            rustfmt.enable = true;
            shellcheck.enable = true;
            shfmt.enable = true;
            statix.enable = true;
          };
        };

        services = {
          postgres = {
            enable = true;
            initialDatabases = [{name = "portfolio";}];
          };
        };
      };
    };
  };
}
