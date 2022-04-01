package main

import (
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	"github.com/kidsan/todo-app/config"
	"github.com/kidsan/todo-app/http"
	"github.com/spf13/cobra"
)

func main() {
	config, err := config.ReadTodoCLIConfig()
	cobra.CheckErr(err)

	client := http.NewClient(fmt.Sprintf("%s:%v", config.Server.Host, config.Server.Port))
	defer client.Close()

	cmd := NewRootCommand(client)

	cobra.CheckErr(cmd.Execute())
}

func NewRootCommand(client todoapp.TodoClient) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "todoctl",
		Short: "todoctl is a cli to manage todos and tasks.",
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(
		NewGetCommand(client),
		NewCreateCommand(client),
		NewUpdateCommand(client),
		NewRemoveCommand(client),
	)

	return rootCmd
}
