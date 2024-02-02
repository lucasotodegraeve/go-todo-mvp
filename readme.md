Building Go modules with nix flakes:
https://nixos.org/manual/nixpkgs/stable/#sec-language-go

vendorHash is set to `null` and `go mod vendor` was run.

Inspect schema
```shell
atlas schema inspect -u "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"
```

Dry-run apply schema
```shell
atlas schema apply -u "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable" --to "file://schema.hcl" --dry-run
```

Starting postgres container
```shell
podman run --rm --name pg -p 5432:5432 -e POSTGRES_PASSWORD=admin -d postgres:16.1
```

Interactive postgres
```shell
psql -h localhost -U postgres
```

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
