package main

import (
	"fmt"
	"strconv"

	todoapp "github.com/kidsan/todo-app"
	"github.com/spf13/cobra"
)

func NewRemoveCommand(client todoapp.TodoClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"remove", "rm"},
		Short:   "delete resources",
	}

	cmd.AddCommand(
		newDeleteTodoCommand(client),
		newDeleteTaskCommand(client),
	)

	return cmd
}

func newDeleteTodoCommand(client todoapp.TodoClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "todo",
		Aliases: []string{"todos"},
		Short:   "delete todos",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cobra.CheckErr(fmt.Errorf("please specify a todo to be deleted"))
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				cobra.CheckErr(fmt.Errorf("invalid id"))
			}

			if id == 0 {
				cobra.CheckErr(fmt.Errorf("invalid id"))
			}

			err = client.DeleteTodo(todoapp.Todo{ID: int32(id)})
			cobra.CheckErr(err)
			Infof("deleted")

		},
	}

	return cmd
}

func newDeleteTaskCommand(client todoapp.TodoClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "task",
		Aliases: []string{"tasks"},
		Short:   "delete tasks",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cobra.CheckErr(fmt.Errorf("please specify a task to be deleted"))
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				cobra.CheckErr(fmt.Errorf("invalid id"))
			}

			if id == 0 {
				cobra.CheckErr(fmt.Errorf("invalid id"))
			}

			err = client.DeleteTask(todoapp.Task{ID: int32(id)})
			cobra.CheckErr(err)
			Infof("deleted")

		},
	}

	return cmd

}
