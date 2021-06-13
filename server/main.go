package main

import (
	"github.com/suradidchao/my-grpc/server/server"
)

func main() {
	server := server.NewBookServer(9000)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
