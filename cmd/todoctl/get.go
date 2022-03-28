package main

import (
	"fmt"
	"strconv"

	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/http"
	"github.com/spf13/cobra"
)

func NewGetCommand(cfg todoapp.CLIConfig) *cobra.Command {
	cmdGet := &cobra.Command{
		Use:     "get",
		Aliases: []string{"show"},
		Short:   "get resources",
	}

	cmdGet.AddCommand(
		newGetTodoCommand(cfg),
	)

	return cmdGet
}

func newGetTodoCommand(cfg todoapp.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "todo",
		Aliases: []string{"todos"},
		Short:   "show todos",
		Run: func(cmd *cobra.Command, args []string) {
			client := http.NewClient(fmt.Sprintf("%s:%v", cfg.Server.Host, cfg.Server.Port))
			defer client.Close()
			if len(args) == 0 {
				todos, err := client.GetAll()

				if err != nil {
					panic(err)
				}

				PrintObject(todos)
				return
			}

			var output []todoapp.Todo

			for _, v := range args {
				id, err := strconv.Atoi(v)
				if err != nil {
					cobra.CheckErr(fmt.Errorf("invalid id: %v", v))
				}
				todo, err := client.Find(todoapp.Todo{ID: int32(id)})
				cobra.CheckErr(err)

				output = append(output, todo)
			}
			PrintObject(output)
		},
	}
	return cmd
}
