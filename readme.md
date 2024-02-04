# Important notes

Building Go modules with nix flakes:
https://nixos.org/manual/nixpkgs/stable/#sec-language-go

vendorHash is set to `null` and `go mod vendor` was run.

BUG: adding items with the same name lead to weird behavior

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

List all databases inside of psql
```shell
\l
```
