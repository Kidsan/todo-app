package adapters

import (
	"context"

	todoapp "github.com/kidsan/todo-app"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (t *TodoRepository) Get(ctx context.Context) ([]todoapp.Todo, error) {
	return []todoapp.Todo{}, nil
}

func (t *TodoRepository) Save(ctx context.Context) ([]todoapp.Todo, error) {
	return []todoapp.Todo{}, nil
}

func (t *TodoRepository) Delete(ctx context.Context) ([]todoapp.Todo, error) {
	return []todoapp.Todo{}, nil
}
