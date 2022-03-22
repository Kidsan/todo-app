package sql

import (
	"context"

	todoapp "github.com/kidsan/todo-app"
	"gorm.io/gorm"
)

// TodoService represents a service for managing todos
type TodoService struct {
	db    *gorm.DB
	todos []todoapp.Todo
}

// NewTodoService returns a new instance of TodoService
func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

// GetTodos returns all todos in the database
func (t *TodoService) GetTodos(ctx context.Context) ([]todoapp.Todo, error) {
	// var result []todoapp.Todo
	// sqlQuery := "select * from todos;"

	// tx := t.db.WithContext(ctx).Raw(sqlQuery).Scan(&result)
	// if tx.Error != nil {
	// 	return nil, fmt.Errorf("adapters: could not list all todos: %w", tx.Error)
	// }
	t.todos = append(t.todos, todoapp.Todo{Name: "foo"})
	return t.todos, nil

	// return result, nil
}
