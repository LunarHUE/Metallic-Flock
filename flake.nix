{
  description = "Compute-Flock Flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      # Systems to support (add more if needed, e.g., aarch64-darwin)
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      
      # Helper function to generate outputs for each system
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      
      # Import nixpkgs for a specific system
      nixpkgsFor = system: import nixpkgs { inherit system; };
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor system;
        in
        {
          default = pkgs.buildGoModule {
            pname = "compute-flock";
            version = "0.1.0";

            # Pulls the repository directly from GitHub
            src = pkgs.fetchFromGitHub {
              owner = "LunarHUE";
              repo = "Compute-Flock";
              rev = "main"; # or a specific commit hash like "a1b2c3d..."
              hash = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="; # START HERE
            };

            # This matches your description: /cmd/compute-flock
            subPackages = [ "cmd/compute-flock" ];

            # Nix needs to verify Go dependencies.
            # Start with a fake hash; Nix will error and give you the real one.
            vendorHash = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="; 
          };
        });

      # Optional: A development shell if you want to run 'go run .' manually
      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor system;
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [ go ];
          };
        });
    };
}
