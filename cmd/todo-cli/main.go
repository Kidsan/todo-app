package main

import (
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/grpc"
)

func main() {
	client := grpc.NewClient("0.0.0.0:3000")

	todos, err := client.GetAll()
	if err != nil {
		panic(err)
	}
	var result []todoapp.Todo

	for _, v := range todos.Todos {
		result = append(result, todoapp.Todo{
			Name:        v.GetName(),
			Description: v.GetDescription(),
		})
	}
	fmt.Println(result)
}
