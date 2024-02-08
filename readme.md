# To-do MVP
Read about it on [my blog](https://lucasotodegraeve.github.io/posts/2024/go-todo-mvp/)

This app constists out of a CLI interface written in Go and a PostgreSQL database to store the items.

The following technologies where used:
- [Go](https://go.dev/)
- [Atlas](https://atlasgo.io/)
- [Nix & Nix Flakes](https://nixos.org/)
- [Podman](https://podman.io/)

## Setup

First build/pull the container image using Nix

```shell
nix build .#postgres
```

Load the image into Docker/Podman. In [Nushell](https://www.nushell.sh/) this looks like

```shell
open result | podman load
```
or in Bash
```shell
podman load < result
```

Start the database using the justfile. This will first start the conainer
and will then use Atlas to apply the schema to the database.

```shell
just postgres
```

The database schema is managed with [Atlas](https://atlasgo.io/) and can be viewed in `schema.hcl`.
There is only a one table with the columns: `id`, `completed` and `name`

To directly interact with the database run the command bellow.
The password for the `postgres` user is `admin`. (You can exit `psql` with ctrl+D)
```shell
just psql
```

To build the Go app
```shell
nix build
```

Run it with
```shell
./result/bin/go-todo-mvp
```

Atlas and psql generate files in the home directory. Clean them up using
```
just clean-global
```

See the `justfile` other commonly used commands.

