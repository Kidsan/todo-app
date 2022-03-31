package main

import (
	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var Verbose bool

func main() {
	config, err := config.ReadTodoCLIConfig()
	cobra.CheckErr(err)

	cmd := NewRootCommand(config)

	cobra.CheckErr(cmd.Execute())
}

func NewRootCommand(cfg todoapp.CLIConfig) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "todoctl",
		Short: "todoctl is a cli to manage todos and tasks.",
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(
		NewGetCommand(cfg),
		NewSaveCommand(cfg),
		NewRemoveCommand(cfg),
	)

	return rootCmd
}

// func main() {
// 	client := http.NewClient("0.0.0.0:3000")
// 	defer client.Close()
// 	// todos, err := client.GetAll()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// fmt.Println(todos)

// 	newTodo := todoapp.Todo{
// 		Name:        "Clean the house",
// 		Description: "Better get to it son",
// 		Tasks: []todoapp.Task{
// 			{Name: "Clean up Task"},
// 			{Name: "Tear down Task"},
// 		},
// 	}

// 	res, _ := client.Save(newTodo)
// 	fmt.Println(res)

// 	toFind := todoapp.Todo{
// 		ID: res.ID,
// 	}
// 	res2, _ := client.Find(toFind)
// 	fmt.Println(res2)

// 	res2.Name = "Updated name"

// 	res3, _ := client.Save(res2)
// 	fmt.Println(res3)
// 	res4 := client.Delete(res3)
// 	fmt.Println("deleted", res4)

// 	todos, err := client.GetAll()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("final", todos)
// }
