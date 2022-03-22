package todoapp

import "context"

type TodoService interface {
	GetTodos(ctx context.Context) ([]Todo, error)
}

type Todo struct {
	Name        string
	Description string
	Tasks       []Task
}
