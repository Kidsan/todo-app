package main

import (
	"fmt"
	"strconv"

	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/http"
	"github.com/spf13/cobra"
)

func NewRemoveCommand(cfg todoapp.CLIConfig) *cobra.Command {
	cmdGet := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"remove", "rm"},
		Short:   "delete resources",
	}

	cmdGet.AddCommand(
		newDeleteTodoCommand(cfg),
	)

	return cmdGet
}

func newDeleteTodoCommand(cfg todoapp.CLIConfig) *cobra.Command {
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
			client := http.NewClient("0.0.0.0:3000")
			defer client.Close()

			err = client.Delete(todoapp.Todo{ID: int32(id)})
			cobra.CheckErr(err)

		},
	}

	return cmd

}
