package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"

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

func addItemDB(db *sql.DB, name string) {
	_, err := db.Query("INSERT INTO items (completed, name) values ($1, $2)", false, name)
	if err != nil {
		log.Fatal(err)
	}
}

func addItem(db *sql.DB) {
	var name string
	input := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Item name").Value(&name),
		),
	)
	err := input.Run()
	if err != nil {
		log.Fatal(nil)
	}
	addItemDB(db, name)
}

func printList(db *sql.DB) {
	list := retreiveItemsDB(db)
	fmt.Println(list)
	// fmt.Printf("[%s] %s", name)

}

func retreiveItemsDB(db *sql.DB) (list []string) {
	rows, err := db.Query("SELECT (name) FROM items")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, name)
	}
	return
}

func cliLoop(db *sql.DB) {
	var selected action

	for {
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

		printList(db)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch selected {
		case new_item:
			addItem(db)
		case rename_item:
		case delete_item:
		case toggle_item:
		}
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
