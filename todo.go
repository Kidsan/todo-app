package todoapp

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v2"
)

type TodoClient interface {
	GetAllTodos() ([]Todo, error)
	FindTodo(Todo) (Todo, error)
	SaveTodo(Todo) (Todo, error)
	DeleteTodo(Todo) error

	DeleteTask(Task) error
	SaveTask(Task) (Task, error)
}

type TodoService interface {
	GetAll(context.Context) ([]Todo, error)
	Find(context.Context, Todo) (Todo, error)
	Update(context.Context, Todo) (Todo, error)
	Delete(context.Context, Todo) error
}

type Todo struct {
	ID          int32
	Name        string
	Description string
	Tasks       []Task `gorm:"ForeignKey:TodoID"`
}

func (t Todo) ToYAML() (string, error) {
	content, err := yaml.Marshal(t)
	if err != nil {
		return "", fmt.Errorf("unable to format yaml %w", err)
	}
	return string(content), nil
}
