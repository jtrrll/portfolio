{
  perSystem = {config, ...}: {
    apps = {
      "portfolio" = {
        program = "${config.packages."portfolio"}/bin/portfolio";
        type = "app";
      };
    };
  };
}
