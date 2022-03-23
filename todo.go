package todoapp

import (
	"context"

	"github.com/google/uuid"
)

type TodoService interface {
	GetAll(context.Context) ([]Todo, error)
	Find(context.Context, Todo) (Todo, error)
	Create(context.Context, Todo) (Todo, error)
	Update(context.Context, Todo) (Todo, error)
	Delete(context.Context, Todo) error
}

type Todo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Complete    bool
	Tasks       []Task `gorm:"ForeignKey:TodoID"`
}
