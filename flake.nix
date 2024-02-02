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
      packages = [
        atlas
        postgresql_16
      ];
    };

    packages."${system}" = {
      default = buildGoModule {
        pname = "Todo-list";
        version = "0.0.1";
        src = self;
        vendorHash = null;
      };
      postgres = dockerTools.pullImage {
        imageName = "postgres";
        imageDigest = "sha256:49c276fa02e3d61bd9b8db81dfb4784fe814f50f778dce5980a03817438293e3";
        sha256 = "1a0qxwy830n2y26wfp37xkjzl7chmz677wk7zvifmjlj0569nc76";
        finalImageName = "postgres";
        finalImageTag = "16.1";      
      };
    };

    
  };
}
