package todoapp

import (
	"context"
)

type TodoService interface {
	GetAll(context.Context) ([]Todo, error)
	Find(context.Context, Todo) (Todo, error)
	Create(context.Context, Todo) (Todo, error)
	Update(context.Context, Todo) (Todo, error)
	Delete(context.Context, Todo) error
}

type Todo struct {
	ID          int
	Name        string
	Description string
	Tasks       []Task `gorm:"ForeignKey:TodoID"`
}
