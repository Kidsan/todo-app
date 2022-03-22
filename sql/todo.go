package sql

import (
	"context"
	"fmt"

	todoapp "github.com/kidsan/todo-app"
)

// TodoService represents a service for managing todos
type TodoService struct {
	db    *DB
	todos []todoapp.Todo
}

// NewTodoService returns a new instance of TodoService
func NewTodoService(db *DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

// GetTodos returns all todos in the database
func (t *TodoService) GetTodos(ctx context.Context) ([]todoapp.Todo, error) {
	var result []todoapp.Todo
	tx := t.db.conn.Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("sql: could not list all todos: %w", tx.Error)
	}
	fmt.Println(result)
	return result, nil
}

// Save creates a new Todo in the database
func (t *TodoService) Save(ctx context.Context, newTodo todoapp.Todo) (todoapp.Todo, error) {
	tx := t.db.conn.Create(&newTodo)
	if tx.Error != nil {
		return todoapp.Todo{}, fmt.Errorf("sql: could not list all todos: %w", tx.Error)
	}
	t.db.conn.Save(&newTodo)

	return newTodo, nil
}
