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
	var todoName string
	var description string
	var tasks []string

	cmd := &cobra.Command{
		Use:     "todo",
		Aliases: []string{"todos"},
		Short:   "create todos",
		Run: func(cmd *cobra.Command, args []string) {
			client := http.NewClient(fmt.Sprintf("%s:%v", cfg.Server.Host, cfg.Server.Port))
			defer client.Close()

			var tasksToCreate []todoapp.Task

			for _, v := range tasks {
				tasksToCreate = append(tasksToCreate, todoapp.Task{Name: v})
			}

			newTodo := todoapp.Todo{
				Name:        todoName,
				Description: description,
				Tasks:       tasksToCreate,
			}

			res, err := client.Save(newTodo)
			cobra.CheckErr(err)
			PrintObject(res)
		},
	}

	cmd.Flags().StringVarP(&todoName, "name", "n", "", "name of the todo item")
	cmd.Flags().StringVarP(&description, "description", "d", "", "description of the todo item")
	cmd.Flags().StringArrayVarP(&tasks, "task", "t", []string{}, "tasks to add to the todo item")

	return cmd

}
