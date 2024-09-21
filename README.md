# portfolio

<!-- markdownlint-disable MD013 -->
![GitHub Actions CI Status](https://img.shields.io/github/actions/workflow/status/jtrrll/portfolio/ci.yaml?branch=main&logo=github&label=CI)
![License](https://img.shields.io/github/license/jtrrll/portfolio?label=License)
<!-- markdownlint-enable MD013 -->

jtrrll's personal portfolio.

## Usage

1. [Install Nix](https://zero-to-nix.com/start/install)
2. Run the following to start all services:

   <!-- markdownlint-disable MD013 -->
   ```sh
   devenv up
   ```
   <!-- markdownlint-enable MD013 -->

3. Run the following to start the `portfolio` application:

   <!-- markdownlint-disable MD013 -->
   ```sh
   nix run .#portfolio
   ```
   <!-- markdownlint-enable MD013 -->
