package main

import (
	"context"
	"log"
	"time"

	"github.com/y-zumi/grpc-go/proto/user"
	"google.golang.org/grpc"
)

func main() {
	// Set up a gRPC client
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUsersClient(conn)

	// Request to gRPC server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FindByID(ctx, &user.FindByIDRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("could not find user: %v", err)
	}
	log.Printf("User: %v", r.User)
}
