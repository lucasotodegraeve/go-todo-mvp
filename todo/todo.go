package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
	"log"
	"strconv"
)

type todoItem struct {
	completed bool
	name      string
}

type todoList struct {
	items []todoItem
}

func (list *todoList) add(item todoItem) {
	list.items = append(list.items, item)
}

func (list *todoList) show() {
	s := "Items:\n"
	for i, item := range list.items {
		var completed string
		if item.completed {
			completed = "X"
		} else {
			completed = " "
		}
		s += fmt.Sprintf("%d. [%s] %s\n", i, completed, item.name)
	}
	fmt.Println(s)
}

func handleAdd(list *todoList) {
	fmt.Print("What is the item of the new item?\n>> ")
	var name string
	fmt.Scanf("%s", &name)

	item := todoItem{completed: false, name: name}
	list.add(item)
}

func getValidIndex(list *todoList) int {
	var input string
	fmt.Print(">> ")
	var index int
	var err error
	for {
		fmt.Scanf("%s", &input)
		index, err = strconv.Atoi(input)
		if err != nil || index < 0 || index >= len(list.items) {
			fmt.Print("Please enter a valid number.\n>> ")
			continue
		}
		break
	}
	return index
}

func handleRemove(list *todoList) {
	fmt.Println("What item would like to delete?")
	list.show()
	index := getValidIndex(list)
	list.items = append(list.items[:index], list.items[index+1:]...)
}

func handleComplete(list *todoList) {
	fmt.Println("What item would you like to mark as completed?")
	list.show()
	index := getValidIndex(list)
	list.items[index].completed = true

}

func cliLoop() {
	list := todoList{}

	for {
		fmt.Println(`
What would you like to do?
[0] Print todos
[1] Add new item
[2] Remove item
[3] Complete item`)

		var option int
		fmt.Printf(">> ")
		fmt.Scanf("%d", &option)

		switch option {
		case 0:
			list.show()
		case 1:
			handleAdd(&list)
		case 2:
			handleRemove(&list)
		case 3:
			handleComplete(&list)
		}
	}

}

func main() {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Hello world")
}
