package todoapp

type Task struct {
	ID     int32
	TodoID int32 `gorm:"column:todo_id" yaml:"-"`
	Name   string
}
