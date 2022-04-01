package sql

import (
	"context"
	"fmt"

	todoapp "github.com/kidsan/todo-app"
)

// taskService represents a service for managing todos
type TaskService struct {
	db    *DB
	todos []todoapp.Task
}

// NewTaskService returns a new instance of TaskService
func NewTaskService(db *DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

// Delete removes a task from the database
func (t *TaskService) Delete(ctx context.Context, toDelete todoapp.Task) error {
	tx := t.db.conn.Delete(&toDelete)
	if tx.Error != nil {
		return fmt.Errorf("sql: could not delete task: %w", tx.Error)
	}

	return nil
}

// Update modifies a Task in the database
func (t *TaskService) Update(ctx context.Context, modifiedTask todoapp.Task) (todoapp.Task, error) {
	tx := t.db.conn.Model(&modifiedTask).Update("name", modifiedTask.Name)
	if tx.Error != nil {
		return todoapp.Task{}, fmt.Errorf("sql: could not update todo: %w", tx.Error)
	}

	return modifiedTask, nil
}
