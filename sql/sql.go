package sql

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	createSchema     = "CREATE SCHEMA IF NOT EXISTS todos;"
	dropTodos        = `DROP TABLE IF EXISTS todos.todos;`
	dropTasks        = `DROP TABLE IF EXISTS todos.tasks;`
	createTodosTable = `CREATE TABLE IF NOT EXISTS todos.todos (
			id SERIAL PRIMARY_KEY,
			name varchar(40),
			description varchar(100),
		);`
	createTasksTable = `
	CREATE TABLE IF NOT EXISTS todos.tasks (
		id SERIAL PRIMARY_KEY,
		todo_id varchar(40),
		name varchar(40),
		CONSTRAINT fk_todo FOREIGN KEY(todo_id) REFERENCES todos.todos(id) ON DELETE CASCADE
	);`
)

// DB represents the database connection.
type DB struct {
	conn *gorm.DB
	DSN  string
}

// NewDB returns a new instance of DB associated with the given datasource name.
func NewDB(dsn string) *DB {
	db := &DB{
		DSN: dsn,
	}
	return db
}

// Open opens the database connection.
func (db *DB) Open() (err error) {
	connection, err := gorm.Open(postgres.Open(db.DSN), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		return fmt.Errorf("api: could not open database: %w", err)
	}

	db.conn = connection

	return nil
}

// migrate sets up the database in the required state.
func (db *DB) RunMigration(ctx context.Context) error {

	migrations := []string{
		createSchema,
		dropTasks,
		dropTodos,
		createTodosTable,
		createTasksTable,
	}

	for _, v := range migrations {
		tx := db.conn.WithContext(ctx).Exec(v)
		if tx.Error != nil {
			return fmt.Errorf("sql: could not create table: %w", tx.Error)
		}
	}
	return nil
}

// Close closes the database connection.
func (db *DB) Close() error {
	// Close database.
	if db.conn != nil {
		connection, err := db.conn.DB()
		if err != nil {
			return err
		}
		return connection.Close()
	}
	return nil
}
