# Important notes

Building Go modules with nix flakes:
https://nixos.org/manual/nixpkgs/stable/#sec-language-go

vendorHash is set to `null` and `go mod vendor` was run.

# Command list
Also see the justfile

Inside `psql` run
```
SELECT current_database();
```
default is `postgres`


Creating a new database
```shell
podman exec -u postgres pg createdb mydb
```
