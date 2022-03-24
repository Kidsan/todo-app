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
func (t *TodoService) GetAll(ctx context.Context) ([]todoapp.Todo, error) {
	var result []todoapp.Todo
	tx := t.db.conn.Preload("Tasks").Find(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("sql: could not list all todos: %w", tx.Error)
	}

	return result, nil
}

// Find returns a specific Todo from the database
func (t *TodoService) Find(ctx context.Context, toFind todoapp.Todo) (todoapp.Todo, error) {
	var result todoapp.Todo
	tx := t.db.conn.Preload("Tasks").Find(&result, toFind.Name)
	if tx.Error != nil {
		return todoapp.Todo{}, fmt.Errorf("sql: could not create todo: %w", tx.Error)
	}

	return result, nil
}

// Create records a new Todo in the database
func (t *TodoService) Create(ctx context.Context, newTodo todoapp.Todo) (todoapp.Todo, error) {
	tx := t.db.conn.Create(&newTodo)
	if tx.Error != nil {
		return todoapp.Todo{}, fmt.Errorf("sql: could not create todo: %w", tx.Error)
	}

	return newTodo, nil
}

// Update modifies a new Todo in the database
func (t *TodoService) Update(ctx context.Context, modifiedTodo todoapp.Todo) (todoapp.Todo, error) {
	tx := t.db.conn.Save(&modifiedTodo)
	if tx.Error != nil {
		return todoapp.Todo{}, fmt.Errorf("sql: could not list update todo: %w", tx.Error)
	}

	return modifiedTodo, nil
}

// Delete removes a Todo and its Tasks from the database
func (t *TodoService) Delete(ctx context.Context, toDelete todoapp.Todo) error {
	tx := t.db.conn.Delete(&toDelete)
	if tx.Error != nil {
		return fmt.Errorf("sql: could not delete todo: %w", tx.Error)
	}

	return nil
}
