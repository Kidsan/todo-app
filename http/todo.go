package http

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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
	todos, err := t.service.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("ports(todos): could not get all todos %w", err)
	}

	r := &pb.TodoListReply{}
	for _, v := range todos {
		var tasks []*pb.TaskReply

		todo := &pb.TodoReply{
			Id:          v.ID.String(),
			Name:        v.Name,
			Description: v.Description,
		}

		for _, j := range v.Tasks {
			tasks = append(tasks, &pb.TaskReply{
				Name: j.Name,
			})
		}

		todo.Tasks = append(todo.Tasks, tasks...)

		r.Todos = append(r.Todos, todo)
	}

	return r, nil
}

func (t TodoGRPCHandler) Save(ctx context.Context, newTodoRequest *pb.TodoRequest) (*pb.TodoReply, error) {
	var tasks []todoapp.Task

	for _, v := range newTodoRequest.GetTasks() {
		tasks = append(tasks, todoapp.Task{
			ID:   uuid.New(),
			Name: v.GetName(),
		})
	}
	newTodo := todoapp.Todo{
		ID:          uuid.New(),
		Name:        newTodoRequest.GetName(),
		Description: newTodoRequest.GetDescription(),
		Tasks:       tasks,
	}
	result, err := t.service.Create(ctx, newTodo)
	if err != nil {
		return &pb.TodoReply{}, fmt.Errorf("ports(todos): could not save new todo %w", err)
	}
	return &pb.TodoReply{Id: result.ID.String(), Name: result.Name, Description: result.Description}, nil

}

func (t TodoGRPCHandler) Find(ctx context.Context, toFind *pb.TodoIdentifier) (*pb.TodoReply, error) {
	id, err := uuid.Parse(toFind.Id)
	result, err := t.service.Find(ctx, todoapp.Todo{
		ID: id,
	})
	if err != nil {
		return &pb.TodoReply{}, fmt.Errorf("ports(todo): could not find todo %w", err)
	}

	var tasks []*pb.TaskReply
	for _, j := range result.Tasks {
		tasks = append(tasks, &pb.TaskReply{
			Id:   j.ID.String(),
			Name: j.Name,
		})
	}

	return &pb.TodoReply{Id: result.ID.String(), Name: result.Name, Description: result.Description, Complete: result.Complete, Tasks: tasks}, nil
}
