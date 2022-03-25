package todoapp

type Task struct {
	ID     int
	TodoID int `gorm:"column:todo_id"`
	Name   string
}
