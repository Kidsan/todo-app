package http

import (
	pb "github.com/kidsan/todo-app/proto"
	"github.com/kidsan/todo-app/todo/adapters"
	"github.com/kidsan/todo-app/todo/domain"
	"github.com/kidsan/todo-app/todo/ports"
)

func (g *GRPCServer) buildTodoServer() pb.TodosServer {
	// collector := adapters.NewAdapterCollector("todo_api")
	// if err := prometheus.Register(collector); err != nil {
	// 	if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
	// 		collector = are.ExistingCollector.(*adapters.Collector)
	// 	} else {
	// 		g.logger.Error("domain: something went wrong while register metrics", zap.Error(err))
	// 	}
	// }
	todoRepository := adapters.NewTodoRepository(g.connection, collector)
	todoService := domain.NewTodoService(todoRepository)
	todoGRPC := ports.NewTodoGRPCHandler(g.logger, todoService)

	return nil
}
