package todoapp

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v2"
)

type TaskService interface {
	Update(context.Context, Task) (Task, error)
	Delete(context.Context, Task) error
}

type Task struct {
	ID     int32
	TodoID int32 `gorm:"column:todo_id" yaml:"-"`
	Name   string
}

func (t Task) ToYAML() (string, error) {
	content, err := yaml.Marshal(t)
	if err != nil {
		return "", fmt.Errorf("unable to format yaml %w", err)
	}
	return string(content), nil
}
