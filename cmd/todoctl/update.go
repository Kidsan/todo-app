package main

import (
	"errors"
	"fmt"
	"strconv"

	todoapp "github.com/kidsan/todo-app"
	"github.com/spf13/cobra"
)

func NewUpdateCommand(client todoapp.TodoClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"update"},
		Short:   "update resources",
	}

	cmd.AddCommand(
		newUpdateTodoCommand(client),
		newUpdateTaskCommand(client),
	)

	return cmd
}

func newUpdateTodoCommand(client todoapp.TodoClient) *cobra.Command {
	var todoName string
	var description string
	var tasks []string
	var id int32

	cmd := &cobra.Command{
		Use:     "todo",
		Aliases: []string{"todos"},
		Short:   "update todos",
		Run: func(cmd *cobra.Command, args []string) {
			var tasksToCreate []todoapp.Task

			for _, v := range tasks {
				tasksToCreate = append(tasksToCreate, todoapp.Task{Name: v})
			}

			if len(args) != 1 {
				cobra.CheckErr(errors.New("Must provide a todo id to update"))
			}

			a, err := strconv.Atoi(args[0])
			if err != nil {
				cobra.CheckErr(fmt.Errorf("invalid todo id: %s", args[0]))
			}
			id = int32(a)

			newTodo := todoapp.Todo{
				ID:          id,
				Name:        todoName,
				Description: description,
				Tasks:       tasksToCreate,
			}

			res, err := client.SaveTodo(newTodo)
			cobra.CheckErr(err)
			PrintObject(cmd, res)
		},
	}

	cmd.Flags().StringVarP(&todoName, "name", "n", "", "name of the todo item")
	cmd.Flags().StringVarP(&description, "description", "d", "", "description of the todo item")
	cmd.Flags().StringArrayVarP(&tasks, "task", "t", []string{}, "tasks to add to the todo item")

	return cmd
}

func newUpdateTaskCommand(client todoapp.TodoClient) *cobra.Command {
	var todoName string
	var id int32

	cmd := &cobra.Command{
		Use:     "task",
		Aliases: []string{"tasks"},
		Short:   "update tasks",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cobra.CheckErr(errors.New("Must provide a task id to update"))
			}

			a, err := strconv.Atoi(args[0])
			if err != nil {
				cobra.CheckErr(fmt.Errorf("invalid task id: %s", args[0]))
			}
			id = int32(a)

			newTask := todoapp.Task{
				ID:   id,
				Name: todoName,
			}

			res, err := client.SaveTask(newTask)
			cobra.CheckErr(err)
			PrintObject(cmd, res)
		},
	}

	cmd.Flags().StringVarP(&todoName, "name", "n", "", "name of the todo item")

	return cmd
}
