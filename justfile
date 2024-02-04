@default:
	just --list

alias w := watch
alias b := build
alias r := run
alias c := clean
alias si := schema-inspect
alias sa := schema-apply

# Build the Go module
build:
	go build src/main.go

# Run the Go module
run:
	go run src/main.go

# Watch `.go` files and recompile
watch:
	watchexec -e go -c clear go run src/main.go

# Remove temporary and generated files
clean:
	rm -f main result

# Inspect the database schema
schema-inspect:
	atlas schema inspect -u "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"

# Apply the database schema
schema-apply apply="false":
	atlas schema apply -u "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable" --to "file://persistance/schema.hcl" {{ if apply == "true" {""} else {"--dry-run"} }}

# Start new postgres container
postgres:
	podman run --rm --name pg -p 5432:5432 --env POSTGRES_PASSWORD=admin --detach postgres:16.1

# List running containers
ps:
	podman ps

# Interactive postgres command line
psql:
	psql -h localhost -U postgres
