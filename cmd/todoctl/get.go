package main

import (
	"fmt"
	"strconv"

	todoapp "github.com/kidsan/todo-app"
	"github.com/spf13/cobra"
)

func NewGetCommand(client todoapp.TodoClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"show"},
		Short:   "get resources",
	}

	cmd.AddCommand(
		newGetTodoCommand(client),
	)

	return cmd
}

func newGetTodoCommand(client todoapp.TodoClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "todo",
		Aliases: []string{"todos"},
		Short:   "show todos",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				todos, err := client.GetAllTodos()

				if err != nil {
					fmt.Println(err)
					panic(err)
				}

				PrintObject(cmd, todos...)
				return
			}

			var output []todoapp.Todo

			for _, v := range args {
				id, err := strconv.Atoi(v)
				if err != nil {
					cobra.CheckErr(fmt.Errorf("invalid id: %v", v))
				}
				todo, err := client.FindTodo(todoapp.Todo{ID: int32(id)})
				cobra.CheckErr(err)

				output = append(output, todo)
			}
			PrintObject(cmd, output...)
		},
	}
	return cmd
}
