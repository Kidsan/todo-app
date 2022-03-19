package ports

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	todosapi "github.com/kidsan/todo-app"
	pb "github.com/kidsan/todo-app/proto"
)

type TodoService interface {
	Get(context.Context) ([]todosapi.Todo, error)
	Save(context.Context, todosapi.Todo) (todosapi.Todo, error)
	Find(context.Context, string) (todosapi.Todo, error)
}

type GRPCHandler struct {
	pb.UnimplementedTodosServer
	service TodoService
	logger  *zap.Logger
}

func NewTodoGRPCHandler(logger *zap.Logger, s TodoService) *GRPCHandler {
	return &GRPCHandler{
		logger:  logger,
		service: s,
	}
}

func (g GRPCHandler) Get(ctx context.Context, _ *pb.GetRequest) (*pb.TodoListReply, error) {
	todos, err := g.service.Get(ctx)
	r := &pb.TodoListReply{}
	for _, v := range todos {
		r.Todos = append(r.Todos, &pb.TodoReply{
			Name: v.Name,
		})
	}
	if err != nil {
		return nil, fmt.Errorf("ports(todos): could not get all todos %w", err)
	}
	return r, nil
}

func (g GRPCHandler) Save(ctx context.Context, newTodoRequest *pb.TodoRequest) (*pb.TodoReply, error) {
	newTodo := todosapi.Todo{
		Name: newTodoRequest.GetName(),
	}
	result, err := g.service.Save(ctx, newTodo)
	if err != nil {
		return &pb.TodoReply{}, fmt.Errorf("ports(todos): could save new todo %w", err)
	}
	return &pb.TodoReply{Name: result.Name}, nil
}

func (g GRPCHandler) Find(ctx context.Context, id *pb.TodoName) (*pb.TodoReply, error) {
	result, err := g.service.Find(ctx, id.GetName())
	if err != nil {
		return &pb.TodoReply{}, fmt.Errorf("ports(todo): could find todo %w", err)
	}
	return &pb.TodoReply{Name: result.Name}, nil
}
