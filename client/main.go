package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/suradidchao/my-grpc/book"
	"google.golang.org/grpc"
)

func main() {
	const port = 9000
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewBookClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	request := pb.Request{
		Id: 12345,
	}

	response, err := client.GetBookById(ctx, &request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
