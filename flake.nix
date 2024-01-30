{
  description = "A very basic flake";

  inputs.nixpkgs.url = "nixpkgs/nixpkgs-unstable";

  outputs = { self, nixpkgs }: 
    let 
      system = "x86_64-linux";
      pkgs = import nixpkgs {
        inherit system;
      };
    in with pkgs; {

    devShells."${system}".default = mkShell {
      nativeBuildInputs = [
        go
      ];
    };

    packages."${system}".default = buildGoModule {
      pname = "Todo-list";
      version = "0.0.1";
      src = self;
      vendorHash = null;
    };
  };
}
