package main

import (
	"github.com/charmbracelet/huh"
	"log"

	"todo-mvp/persistance"

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

// type todoItem struct {
// 	id        int
// 	completed bool
// 	name      string
// }

// type todoList struct {
// 	items []todoItem
// }

// func (list *todoList) add(item todoItem) {
// 	list.items = append(list.items, item)
// }

// func addItem(db *sql.DB) {
// }

// Adds an item to the postgres database.
// By default the item is not marked as completed an is called name.
func addItemDB(db *sql.DB, name string) {
	_, err := db.Query("INSERT INTO items (completed, name) values ($1, $2)", false, name)
	if err != nil {
		log.Fatal(err)
	}
}

// func getValidIndex(list *todoList) int {
// 	var input string
// 	fmt.Print(">> ")
// 	var index int
// 	var err error
// 	for {
// 		fmt.Scanf("%s", &input)
// 		index, err = strconv.Atoi(input)
// 		if err != nil || index < 0 || index >= len(list.items) {
// 			fmt.Print("Please enter a valid number.\n>> ")
// 			continue
// 		}
// 		break
// 	}
// 	return index
// }

// func handleRemove(list *todoList) {
// 	fmt.Println("What item would like to delete?")
// 	list.show()
// 	index := getValidIndex(list)
// 	list.items = append(list.items[:index], list.items[index+1:]...)
// }

// func handleComplete(list *todoList) {
// 	fmt.Println("What item would you like to mark as completed?")
// 	list.show()
// 	index := getValidIndex(list)
// 	list.items[index].completed = true

// }

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

// func init_list(db *sql.DB) todoList {
// 	list := todoList{}
// 	rows, err := db.Query("SELECT * FROM items")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id int
// 		var completed bool
// 		var name string

// 		if err := rows.Scan(&id, &completed, &name); err != nil {
// 			log.Fatal(err)
// 		}

// 		item := todoItem{id, completed, name}
// 		list.add(item)
// 	}
// 	return list
// }

func main() {

	db := connect()
	// model := model{
	// 	list: init_list(db),
	// 	db:   db,
	// }
	cliLoop(db)
}
