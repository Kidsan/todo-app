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

func (c *Client) GetAll() ([]todoapp.Todo, error) {
	todos, err := c.grpc.Get(context.Background(), &pb.GetRequest{})
	if err != nil {
		return nil, err
	}

	var result []todoapp.Todo
	for _, v := range todos.GetTodos() {
		var tasks []todoapp.Task

		for _, v := range v.Tasks {
			tasks = append(tasks, todoapp.Task{
				Name: v.Name,
			})
		}
		if err != nil {
			continue
		}
		result = append(result, todoapp.Todo{
			Name:        v.Name,
			Description: v.Description,
			Tasks:       tasks,
		})
	}
	return result, nil
}

func (c *Client) Save(newTodo todoapp.Todo) (todoapp.Todo, error) {
	var tasks []*pb.TaskRequest

	for _, v := range newTodo.Tasks {
		tasks = append(tasks, &pb.TaskRequest{
			Name: v.Name,
		})
	}
	todo, err := c.grpc.Save(context.Background(), &pb.TodoRequest{
		Name:        newTodo.Name,
		Description: newTodo.Description,
		Tasks:       tasks,
	})

	if err != nil {
		return todoapp.Todo{}, err
	}

	var savedTasks []todoapp.Task

	for _, v := range savedTasks {
		savedTasks = append(savedTasks, todoapp.Task{
			Name: v.Name,
		})
	}

	return todoapp.Todo{
		Name:        todo.Name,
		Description: todo.Description,
		Tasks:       savedTasks,
	}, nil
}

func (c *Client) Find(toFind todoapp.Todo) (todoapp.Todo, error) {
	todo, err := c.grpc.Find(context.Background(), &pb.TodoIdentifier{Id: toFind.ID.String()})
	if err != nil {
		return todoapp.Todo{}, err
	}

	result := todoapp.Todo{
		Name:        todo.Name,
		Description: todo.Description,
	}
	for _, v := range todo.Tasks {
		result.Tasks = append(result.Tasks, todoapp.Task{
			Name: v.Name,
		})
	}
	return result, nil
}

func (c *Client) Update(toUpdate todoapp.Todo) (todoapp.Todo, error) {
	var tasks []*pb.TaskRequest

	for _, v := range toUpdate.Tasks {
		tasks = append(tasks, &pb.TaskRequest{
			Name: v.Name,
		})
	}
	todo, err := c.grpc.Save(context.Background(), &pb.TodoRequest{
		Name:        toUpdate.Name,
		Description: toUpdate.Description,
		Tasks:       tasks,
	})

	if err != nil {
		return todoapp.Todo{}, err
	}

	var savedTasks []todoapp.Task

	for _, v := range savedTasks {
		savedTasks = append(savedTasks, todoapp.Task{
			Name: v.Name,
		})
	}

	return todoapp.Todo{
		Name:        todo.Name,
		Description: todo.Description,
		Tasks:       savedTasks,
	}, nil
}
