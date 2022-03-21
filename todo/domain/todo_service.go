package domain

import (
	"context"

	todoapp "github.com/kidsan/todo-app"
)

type TodoRepository interface{}

type TodoService struct{}

func NewTodoService(TodoRepository) *TodoService {
	return &TodoService{}
}

func (t *TodoService) Get(ctx context.Context) ([]todoapp.Todo, error) {
	return []todoapp.Todo{}, nil
}

func (t *TodoService) Save(ctx context.Context) ([]todoapp.Todo, error) {
	return []todoapp.Todo{}, nil
}

func (t *TodoService) Find(ctx context.Context) (todoapp.Todo, error) {
	return todoapp.Todo{}, nil
}
