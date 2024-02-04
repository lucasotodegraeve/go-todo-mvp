package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"

	"database/sql"

	_ "github.com/lib/pq"
)

type todoItem struct {
	id        int
	completed bool
	name      string
}

type itemAction int

const (
	delete itemAction = iota
	rename
	toggle
	noop
)

func simpleUserInput(title string, placeholder string) string {
	s := placeholder
	input := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title(title).Value(&s),
		),
	)
	err := input.Run()
	if err != nil {
		log.Fatal(nil)
	}
	return s
}

func addItemDB(db *sql.DB, name string) {
	_, err := db.Query("INSERT INTO items (completed, name) values ($1, $2)", false, name)
	if err != nil {
		log.Fatal(err)
	}
}

func addItem(db *sql.DB) {
	name := simpleUserInput("Item name", "")
	addItemDB(db, name)
}

func retreiveItemsDB(db *sql.DB) (list []todoItem) {
	rows, err := db.Query("SELECT * FROM items ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var completed bool
		var id int
		err := rows.Scan(&id, &completed, &name)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, todoItem{id, completed, name})
	}
	return
}

func retreiveItemWithIdDB(db *sql.DB, id int) (item todoItem) {
	rows, err := db.Query("SELECT * FROM items WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		// DRY todo: add reflection
		var name string
		var completed bool
		var id int
		err := rows.Scan(&id, &completed, &name)
		if err != nil {
			log.Fatal(err)
		}
		item = todoItem{id, completed, name}
	}
	return
}

func deleteItemWithIdDB(db *sql.DB, id int) {
	_, err := db.Query("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func renameItemWithIdDB(db *sql.DB, id int, newName string) {
	_, err := db.Query("UPDATE items SET name = $2 WHERE id = $1", id, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func renameItem(db *sql.DB, id int) {
	item := retreiveItemWithIdDB(db, id)
	newName := simpleUserInput("New item name", item.name)
	renameItemWithIdDB(db, id, newName)
}

func toggleItemWithIdDB(db *sql.DB, id int) {
	item := retreiveItemWithIdDB(db, id)
	_, err := db.Query("UPDATE items SET completed = $2 WHERE id = $1", id, !item.completed)
	if err != nil {
		log.Fatal(err)
	}
}

func itemOptions(db *sql.DB, id int) {
	var selected itemAction
	item := retreiveItemWithIdDB(db, id)
	s := formatItem(item)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[itemAction]().Title(s).Options(
				huh.NewOption("Rename", rename),
				huh.NewOption("Delete", delete),
				huh.NewOption("Toggle", toggle),
				huh.NewOption("Back", noop),
			).Value(&selected),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch selected {
	case delete:
		deleteItemWithIdDB(db, item.id)
	case rename:
		renameItem(db, item.id)
	case toggle:
		toggleItemWithIdDB(db, item.id)
	}
}

func formatItem(item todoItem) string {
	var cross string
	if item.completed {
		cross = "X"
	} else {
		cross = " "
	}
	return fmt.Sprintf("[%s] %s", cross, item.name)
}

func cliLoop(db *sql.DB) {
	var selected int

	for {

		list := retreiveItemsDB(db)
		var options = []huh.Option[int]{huh.NewOption("Add new item", -1)}
		for _, item := range list {
			s := formatItem(item)
			options = append(options, huh.NewOption(s, item.id))
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().Title("Select or add item").Options(options...).Value(&selected),
			),
		)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		if selected < 0 {
			addItem(db)
		} else {
			itemOptions(db, selected)
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
