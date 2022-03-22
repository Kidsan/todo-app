package grpc

import (
	"context"
	"log"

	pb "github.com/kidsan/todo-app/proto"
	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	grpc pb.TodosClient
}

func NewClient(URL string) *Client {
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	c := pb.NewTodosClient(conn)

	return &Client{
		conn: conn,
		grpc: c,
	}
}

func (c *Client) GetAll() (*pb.TodoListReply, error) {
	return c.grpc.Get(context.Background(), &pb.GetRequest{})
}
