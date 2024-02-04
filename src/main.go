package main

import (
	"github.com/charmbracelet/huh"
	"log"

	// "todo-mvp/persistance"

	"database/sql"
	_ "github.com/lib/pq"
)

type action int

const (
	new_item action = iota
	rename_item
	delete_item
	toggle_item
)

// Adds an item to the postgres database.
// By default the item is not marked as completed an is called name.
func addItemDB(db *sql.DB, name string) {
	_, err := db.Query("INSERT INTO items (completed, name) values ($1, $2)", false, name)
	if err != nil {
		log.Fatal(err)
	}
}

func cliLoop(db *sql.DB) {
	var selected action
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[action]().Title("Select action").Options(
				huh.NewOption("New Item", new_item),
				huh.NewOption("Toggle completion", toggle_item),
				huh.NewOption("Rename item", rename_item),
				huh.NewOption("Delete item", delete_item),
			).Value(&selected),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch selected {
	case new_item:
	case rename_item:
	case delete_item:
	case toggle_item:
	}

}

func connect() *sql.DB {
	connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {

	db := connect()
	cliLoop(db)
}
