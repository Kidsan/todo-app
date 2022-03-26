package main

import (
	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/http"
	"github.com/spf13/cobra"
)

func NewSaveCommand(cfg todoapp.CLIConfig) *cobra.Command {
	cmdGet := &cobra.Command{
		Use:     "save",
		Aliases: []string{"show"},
		Short:   "save resources",
	}

	cmdGet.AddCommand(
		newSaveTodoCommand(cfg),
	)

	return cmdGet
}

func newSaveTodoCommand(cfg todoapp.CLIConfig) *cobra.Command {
	// var name string
	// var serviceName string

	cmd := &cobra.Command{
		Use:     "todo",
		Aliases: []string{"todos"},
		Short:   "show todos",
		Run: func(cmd *cobra.Command, args []string) {
			client := http.NewClient("0.0.0.0:3000")
			defer client.Close()

			newTodo := todoapp.Todo{
				Name:        "Clean the house",
				Description: "Better get to it son",
				Tasks: []todoapp.Task{
					{Name: "Clean up Task"},
					{Name: "Tear down Task"},
				},
			}

			res, _ := client.Save(newTodo)
			PrintObject(res)

		},
	}
	// cmd.Flags().StringVarP(&name, "name", "", "", "get route via name")
	// cmd.Flags().StringVarP(&serviceName, "serviceName", "", "", "get multiple routes via service name (host)")

	return cmd

}
