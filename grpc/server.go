package grpc

import (
	"fmt"
	"log"
	"net"

	todoapp "github.com/kidsan/todo-app"
	pb "github.com/kidsan/todo-app/proto"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	logger *zap.Logger
	server *grpc.Server
	config todoapp.ServerConfig

	todoService todoapp.TodoService
}

func NewGRPCServer(config todoapp.Config, logger *zap.Logger, todoService todoapp.TodoService) *GRPCServer {
	grpcServer := &GRPCServer{
		config:      config.Server,
		logger:      logger,
		server:      grpc.NewServer(),
		todoService: todoService,
	}
	pb.RegisterTodosServer(grpcServer.server, grpcServer.buildTodoServer())
	return grpcServer
}

func (s *GRPCServer) Start() {
	s.logger.Info(fmt.Sprintf("GRPC Application listening on port %d", s.config.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.config.Port))
	if err != nil {
		s.logger.Sugar().Fatalf("failed to listen: %v", err)
	}
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
