package main

import (
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/http"
)

func main() {
	client := http.NewClient("0.0.0.0:3000")
	defer client.Close()
	// todos, err := client.GetAll()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(todos)

	newTodo := todoapp.Todo{
		Name:        "Clean the house",
		Description: "Better get to it son",
		Tasks: []todoapp.Task{
			{Name: "Clean up Task"},
		},
	}

	res, _ := client.Save(newTodo)
	fmt.Println(res)

	toFind := todoapp.Todo{
		Name: res.Name,
	}
	res2, _ := client.Find(toFind)
	fmt.Println(res2)
}
