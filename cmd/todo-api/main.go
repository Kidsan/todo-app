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
	server := http.NewServer(config, logger, todoService)

	server.Start()
}

// func runMigration(config todosapi.DatabaseConfig) error {
// 	migration, err := migration.NewMigration(config)
// 	if err != nil {
// 		return err
// 	}
// 	defer migration.Close()

// 	return migration.Up()
// }

// func openDBConnection(dsn, databaseName string) (*gorm.DB, error) {
// 	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
// 	})

// 	if err != nil {
// 		return nil, fmt.Errorf("api: could not open database: %w", err)
// 	}

// 	return connection, nil
// }
