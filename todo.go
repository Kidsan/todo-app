package todoapp

import (
	"context"

	"github.com/google/uuid"
)

type TodoService interface {
	GetTodos(ctx context.Context) ([]Todo, error)
	Save(ctx context.Context, newTodo Todo) (Todo, error)
}

type Todo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Tasks       []Task
}
