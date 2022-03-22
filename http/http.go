package http

import (
	"context"
	"log"

	"github.com/google/uuid"
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

func (c *Client) GetAll() ([]todoapp.Todo, error) {
	todos, err := c.grpc.Get(context.Background(), &pb.GetRequest{})
	if err != nil {
		return nil, err
	}

	var result []todoapp.Todo
	for _, v := range todos.GetTodos() {
		id, err := uuid.Parse(v.Id)
		if err != nil {
			continue
		}
		result = append(result, todoapp.Todo{
			ID:          id,
			Name:        v.Name,
			Description: v.Description,
			Tasks:       []todoapp.Task{},
		})
	}
	return result, nil
}

func (c *Client) Save(newTodo todoapp.Todo) (todoapp.Todo, error) {
	todo, err := c.grpc.Save(context.Background(), &pb.TodoRequest{
		Name:        newTodo.Name,
		Description: newTodo.Description,
	})
	if err != nil {
		return todoapp.Todo{}, err
	}

	id, err := uuid.Parse(todo.Id)
	if err != nil {
		return todoapp.Todo{}, err
	}

	return todoapp.Todo{
		ID:          id,
		Name:        todo.Name,
		Description: todo.Description,
	}, nil
}
