{
  description = "go flake";

  inputs = {
      nixpkgs.url = "nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs, ... } @ inputs:
	let
	system = "x86_64-linux";
	pkgs = import inputs.nixpkgs {inherit system;};
	in {
	    devShells.${system}.go = pkgs.mkShell {
		buildInputs = with pkgs; [
			go 
			gcc
			libcap
		];

	    shellHook = ''
			kitty
	    	echo "Entered Go Dev environment"
		'';
	};
  };
}
