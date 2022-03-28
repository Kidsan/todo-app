package main

import (
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/http"
	"github.com/spf13/cobra"
)

func NewSaveCommand(cfg todoapp.CLIConfig) *cobra.Command {
	cmdGet := &cobra.Command{
		Use:     "save",
		Aliases: []string{"save", "create", "update", "apply"},
		Short:   "save resources",
	}

	cmdGet.AddCommand(
		newSaveTodoCommand(cfg),
	)

	return cmdGet
}

func newSaveTodoCommand(cfg todoapp.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "todo",
		Aliases: []string{"todos"},
		Short:   "show todos",
		Run: func(cmd *cobra.Command, args []string) {
			client := http.NewClient(fmt.Sprintf("%s:%v", cfg.Server.Host, cfg.Server.Port))
			defer client.Close()

			newTodo := todoapp.Todo{
				Name:        "Clean the house",
				Description: "Better get to it son!",
				Tasks: []todoapp.Task{
					{Name: "Clean up Task"},
					{Name: "Tear down Task"},
				},
			}

			res, err := client.Save(newTodo)
			cobra.CheckErr(err)
			PrintObject(res)
		},
	}

	return cmd

}
