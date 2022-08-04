package main

import (
	"context"
	"log"
	"os"

	ps "../genpassword/genpassword"
	"google.golang.org/grpc"
)

func main() {

	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())

	client := ps.NewPasswordGeneratorServiceClient(conn)

	sample := os.Args[1]

	resp, err := client.Generate(context.Background(),
		&ps.PasswordRequest{Sample: sample})

	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}
	log.Println("New password:", resp.Password)
}
