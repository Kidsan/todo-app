package http

import (
	"context"
	"log"

	todoapp "github.com/kidsan/todo-app"
	pb "github.com/kidsan/todo-app/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	grpc pb.TodosClient
}

// NewClient returns a client instance.
// Client instances should have Close() called on them when
// finished.
func NewClient(URL string) *Client {
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewTodosClient(conn)

	return &Client{
		conn: conn,
		grpc: c,
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetAllTodos() ([]todoapp.Todo, error) {
	todos, err := c.grpc.Get(context.Background(), &pb.GetRequest{})
	if err != nil {
		return nil, err
	}

	var result []todoapp.Todo
	for _, v := range todos.GetTodos() {
		var tasks []todoapp.Task

		for _, v := range v.Tasks {
			tasks = append(tasks, todoapp.Task{
				ID:     v.Id,
				TodoID: v.TodoId,
				Name:   v.Name,
			})
		}
		if err != nil {
			continue
		}
		result = append(result, todoapp.Todo{
			ID:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Tasks:       tasks,
		})
	}
	return result, nil
}

func (c *Client) SaveTodo(newTodo todoapp.Todo) (todoapp.Todo, error) {
	var tasks []*pb.Task

	for _, v := range newTodo.Tasks {
		tasks = append(tasks, &pb.Task{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	todo, err := c.grpc.Save(context.Background(), &pb.Todo{
		Id:          newTodo.ID,
		Name:        newTodo.Name,
		Description: newTodo.Description,
		Tasks:       tasks,
	})

	if err != nil {
		return todoapp.Todo{}, err
	}

	var savedTasks []todoapp.Task

	for _, v := range todo.Tasks {
		savedTasks = append(savedTasks, todoapp.Task{
			ID:     v.Id,
			TodoID: v.TodoId,
			Name:   v.Name,
		})
	}

	return todoapp.Todo{
		ID:          todo.Id,
		Name:        todo.Name,
		Description: todo.Description,
		Tasks:       savedTasks,
	}, nil
}

func (c *Client) FindTodo(toFind todoapp.Todo) (todoapp.Todo, error) {
	todo, err := c.grpc.Find(context.Background(), &pb.Todo{Id: toFind.ID})
	if err != nil {
		return todoapp.Todo{}, err
	}

	result := todoapp.Todo{
		ID:          todo.Id,
		Name:        todo.Name,
		Description: todo.Description,
	}
	for _, v := range todo.Tasks {
		result.Tasks = append(result.Tasks, todoapp.Task{
			ID:     v.Id,
			TodoID: v.TodoId,
			Name:   v.Name,
		})
	}
	return result, nil
}

func (c *Client) DeleteTodo(toDelete todoapp.Todo) error {
	_, err := c.grpc.Delete(context.Background(), &pb.Todo{Id: toDelete.ID})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteTask(toDelete todoapp.Task) error {
	_, err := c.grpc.DeleteTask(context.Background(), &pb.Task{Id: toDelete.ID})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SaveTask(newTask todoapp.Task) (todoapp.Task, error) {
	task, err := c.grpc.SaveTask(context.Background(), &pb.Task{
		Id:   newTask.ID,
		Name: newTask.Name,
	})

	if err != nil {
		return todoapp.Task{}, err
	}

	return todoapp.Task{ID: task.Id, Name: task.Name}, nil
}
