package todoapp

import (
	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID
	TodoID uuid.UUID
	Name   string
}
