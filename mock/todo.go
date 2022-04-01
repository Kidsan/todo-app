package mock

import (
	"fmt"

	todoapp "github.com/kidsan/todo-app"
)

type TodoClient struct {
	todos     []todoapp.Todo
	todoIndex int32
}

func NewTodoClient(URL string) *TodoClient {
	return &TodoClient{todoIndex: 1}
}

func (t *TodoClient) Close() error {
	return nil
}

func (t *TodoClient) GetAllTodos() ([]todoapp.Todo, error) {
	return t.todos, nil
}

func (t *TodoClient) SaveTodo(newTodo todoapp.Todo) (todoapp.Todo, error) {
	newTodo.ID = t.todoIndex
	t.todoIndex += 1
	t.todos = append(t.todos, newTodo)
	return newTodo, nil
}

func (t *TodoClient) FindTodo(toFind todoapp.Todo) (todoapp.Todo, error) {

	for _, v := range t.todos {
		if toFind.ID == v.ID {
			return v, nil
		}
	}

	return todoapp.Todo{}, fmt.Errorf("could not find todo %d", toFind.ID)
}

func (t *TodoClient) DeleteTodo(toDelete todoapp.Todo) error {
	return nil
}

func (t *TodoClient) DeleteTask(toDelete todoapp.Task) error {

	return nil
}

func (t *TodoClient) SaveTask(newTask todoapp.Task) (todoapp.Task, error) {

	return todoapp.Task{ID: 1, Name: "Task Name"}, nil
}
