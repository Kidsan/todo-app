package http

import (
	"fmt"
	"log"
	"net"

	todoapp "github.com/kidsan/todo-app"
	pb "github.com/kidsan/todo-app/proto"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	logger     *zap.Logger
	server     *grpc.Server
	connection *gorm.DB
	config     todoapp.ServerConfig
}

func NewGRPCServer(config todoapp.Config, logger *zap.Logger, connection *gorm.DB) *GRPCServer {
	grpcServer := &GRPCServer{
		config:     config.Server,
		logger:     logger,
		connection: connection,
		server:     grpc.NewServer(),
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
