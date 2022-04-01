package main

import (
	"context"

	"github.com/kidsan/todo-app/config"
	"github.com/kidsan/todo-app/http"
	"github.com/kidsan/todo-app/logger"
	"github.com/kidsan/todo-app/sql"
)

func main() {
	logger := logger.NewLogger()
	config, err := config.Read()
	if err != nil {
		panic(err)
	}

	database := sql.NewDB(config.Database.DSN())
	if err := database.Open(); err != nil {
		panic(err)
	}
	defer database.Close()

	if err := database.RunMigration(context.Background()); err != nil {
		panic(err)
	}
	todoService := sql.NewTodoService(database)
	taskService := sql.NewTaskService(database)
	server := http.NewServer(config, logger, todoService, taskService)

	server.Start()
}
