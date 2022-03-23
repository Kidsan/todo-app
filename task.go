package todoapp

import (
	"github.com/google/uuid"
)

type Task struct {
	ID       uuid.UUID
	TodoID   uuid.UUID `gorm:"column:todo_id"`
	Name     string
	Complete bool
}
