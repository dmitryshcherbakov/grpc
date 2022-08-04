package main

import (
	"context"
	"log"
	"net"

	//garbler "github.com/michaelbironneau/garbler/lib"
	ps "./proto"

	"google.golang.org/grpc"
)

type PasswordGeneratorServiceServer struct {
}

func (s *PasswordGeneratorServiceServer) Generate(ctx context.Context,
	req *ps.PasswordRequest) (*ps.PasswordResponse, error) {

	var err error
	response := new(ps.PasswordResponse)

	//requirements := garbler.MakeRequirements(req.Sample)
	response.Password = "ento_super_password"

	return response, err
}

func main() {
	server := grpc.NewServer()

	instance := new(PasswordGeneratorServiceServer)

	ps.RegisterPasswordGeneratorServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
