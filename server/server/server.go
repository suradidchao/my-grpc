package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/suradidchao/my-grpc/book"
	"google.golang.org/grpc"
)

type BookServer struct {
	pb.UnimplementedBookServer
	Port int
}

func (s *BookServer) GetBookById(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println("Book id: ", req.Id)
	response := pb.Response{
		Id:     123,
		Title:  "The subtle art of not giving a fuck",
		Author: "Mark Manson",
		Pages:  450,
	}
	return &response, nil
}

func (s *BookServer) Start() error {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookServer(grpcServer, s)

	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func NewBookServer(port int) BookServer {
	return BookServer{
		Port: port,
	}
}
