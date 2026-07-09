{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-26.05.tar.gz") {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.rustup
    pkgs.mdbook
    pkgs.mdbook-d2
    pkgs.d2
    pkgs.dprint
  ];
  shellHook = ''
    export PATH=$PATH:.cargo/bin
  '';
}

