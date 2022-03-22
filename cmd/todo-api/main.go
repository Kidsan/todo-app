package main

import (
	"github.com/kidsan/todo-app/config"
	"github.com/kidsan/todo-app/grpc"
	"github.com/kidsan/todo-app/logger"
	"github.com/kidsan/todo-app/sql"
)

type PortListener interface {
	Start()
}

func main() {
	logger := logger.NewLogger()
	config, err := config.Read()
	if err != nil {
		panic(err)
	}

	// // if err := runMigration(config.Database); err != nil {
	// // 	panic(err)
	// // }

	// dbConnection, err := openDBConnection(config.Database.DSN(), config.Database.Database)
	// if err != nil {
	// 	panic(err)
	// }

	todoService := sql.NewTodoService(nil)
	server := grpc.NewServer(config, logger, todoService)

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
