package http

import (
	"fmt"
	"log"
	"net"

	todoapp "github.com/kidsan/todo-app"
	pb "github.com/kidsan/todo-app/proto"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

type Server struct {
	logger     *zap.Logger
	grpcServer *grpc.Server
	config     todoapp.ServerConfig

	todoService todoapp.TodoService
	taskService todoapp.TaskService
}

func NewServer(config todoapp.APIConfig, logger *zap.Logger, todoService todoapp.TodoService, taskService todoapp.TaskService) *Server {
	server := &Server{
		config:      config.Server,
		logger:      logger,
		grpcServer:  grpc.NewServer(),
		todoService: todoService,
		taskService: taskService,
	}
	pb.RegisterTodosServer(server.grpcServer, server.buildTodoServer())
	return server
}

func (s *Server) Start() {
	s.logger.Info(fmt.Sprintf("GRPC Application listening on port %d", s.config.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		s.logger.Sugar().Fatalf("failed to listen: %v", err)
	}
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
