package grpc

import (
	"context"
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	pb "github.com/kidsan/todo-app/proto"
	"go.uber.org/zap"
)

type TodoGRPCHandler struct {
	pb.UnimplementedTodosServer
	service todoapp.TodoService
	logger  *zap.Logger
}

func (g *Server) buildTodoServer() pb.TodosServer {
	return TodoGRPCHandler{
		service: g.todoService,
		logger:  g.logger,
	}
}

func (t TodoGRPCHandler) Get(ctx context.Context, _ *pb.GetRequest) (*pb.TodoListReply, error) {
	todos, err := t.service.GetTodos(ctx)
	if err != nil {
		return nil, fmt.Errorf("ports(todos): could not get all todos %w", err)
	}

	r := &pb.TodoListReply{}
	for _, v := range todos {
		r.Todos = append(r.Todos, &pb.TodoReply{
			Name: v.Name,
		})
	}

	return r, nil
}

func (t TodoGRPCHandler) Save(ctx context.Context, newTodoRequest *pb.TodoRequest) (*pb.TodoReply, error) {
	// newTodo := todosapi.Todo{
	// 	Name: newTodoRequest.GetName(),
	// }
	// result, err := g.service.Save(ctx, newTodo)
	// if err != nil {
	// 	return &pb.TodoReply{}, fmt.Errorf("ports(todos): could save new todo %w", err)
	// }
	// return &pb.TodoReply{Name: result.Name}, nil
	return nil, nil
}

func (t TodoGRPCHandler) Find(ctx context.Context, id *pb.TodoName) (*pb.TodoReply, error) {
	// result, err := g.service.Find(ctx, id.GetName())
	// if err != nil {
	// 	return &pb.TodoReply{}, fmt.Errorf("ports(todo): could find todo %w", err)
	// }
	// return &pb.TodoReply{Name: result.Name}, nil
	return nil, nil
}
