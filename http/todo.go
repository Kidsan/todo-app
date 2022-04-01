package http

import (
	"context"
	"fmt"

	todoapp "github.com/kidsan/todo-app"
	pb "github.com/kidsan/todo-app/proto"
	"go.uber.org/zap"
)

type TodoGRPCHandler struct {
	pb.UnimplementedTodosServer
	todoService todoapp.TodoService
	taskService todoapp.TaskService
	logger      *zap.Logger
}

func (g *Server) buildTodoServer() pb.TodosServer {
	return TodoGRPCHandler{
		todoService: g.todoService,
		taskService: g.taskService,
		logger:      g.logger,
	}
}

func (t TodoGRPCHandler) Get(ctx context.Context, _ *pb.GetRequest) (*pb.TodoList, error) {
	todos, err := t.todoService.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get all todos %w", err)
	}

	r := &pb.TodoList{}
	for _, v := range todos {
		var tasks []*pb.Task

		for _, j := range v.Tasks {
			tasks = append(tasks, &pb.Task{
				Id:     j.ID,
				TodoId: j.TodoID,
				Name:   j.Name,
			})
		}

		todo := &pb.Todo{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Tasks:       tasks,
		}

		r.Todos = append(r.Todos, todo)
	}

	return r, nil
}

func (t TodoGRPCHandler) Save(ctx context.Context, newTodoRequest *pb.Todo) (*pb.Todo, error) {
	var tasks []todoapp.Task

	for _, v := range newTodoRequest.GetTasks() {
		tasks = append(tasks, todoapp.Task{
			ID:   v.GetId(),
			Name: v.GetName(),
		})
	}
	newTodo := todoapp.Todo{
		ID:          newTodoRequest.GetId(),
		Name:        newTodoRequest.GetName(),
		Description: newTodoRequest.GetDescription(),
		Tasks:       tasks,
	}
	result, err := t.todoService.Update(ctx, newTodo)
	if err != nil {
		return &pb.Todo{}, fmt.Errorf("could not save new todo %w", err)
	}

	var savedTasks []*pb.Task
	for _, j := range result.Tasks {
		savedTasks = append(savedTasks, &pb.Task{
			Id:     j.ID,
			TodoId: j.TodoID,
			Name:   j.Name,
		})
	}

	return &pb.Todo{Id: result.ID, Name: result.Name, Description: result.Description, Tasks: savedTasks}, nil
}

func (t TodoGRPCHandler) Find(ctx context.Context, toFind *pb.Todo) (*pb.Todo, error) {
	result, err := t.todoService.Find(ctx, todoapp.Todo{
		ID: toFind.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("could not find todo")
	}

	var tasks []*pb.Task
	for _, j := range result.Tasks {
		tasks = append(tasks, &pb.Task{
			Id:     j.ID,
			TodoId: j.TodoID,
			Name:   j.Name,
		})
	}

	return &pb.Todo{Id: result.ID, Name: result.Name, Description: result.Description, Tasks: tasks}, nil
}

func (t TodoGRPCHandler) Delete(ctx context.Context, toDelete *pb.Todo) (*pb.GetRequest, error) {
	err := t.todoService.Delete(ctx, todoapp.Todo{
		ID: toDelete.Id,
	})
	if err != nil {
		return &pb.GetRequest{}, fmt.Errorf("could not find todo %w", err)
	}

	return &pb.GetRequest{}, nil
}

func (t TodoGRPCHandler) DeleteTask(ctx context.Context, toDelete *pb.Task) (*pb.GetRequest, error) {
	err := t.taskService.Delete(ctx, todoapp.Task{
		ID: toDelete.Id,
	})
	if err != nil {
		return &pb.GetRequest{}, fmt.Errorf("could not delete Task %w", err)
	}

	return &pb.GetRequest{}, nil
}

func (t TodoGRPCHandler) SaveTask(ctx context.Context, newTaskRequest *pb.Task) (*pb.Task, error) {
	newTask := todoapp.Task{
		ID:   newTaskRequest.GetId(),
		Name: newTaskRequest.GetName(),
	}
	result, err := t.taskService.Update(ctx, newTask)
	if err != nil {
		return &pb.Task{}, fmt.Errorf("could not save task %w", err)
	}

	return &pb.Task{Id: result.ID, Name: result.Name}, nil
}
