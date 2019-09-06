package main

import (
	"context"
	"log"
	"time"

	pb "github.com/y-zumi/grpc-go/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a gRPC client
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUsersClient(conn)

	// Request to gRPC server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FindByID(ctx, &pb.FindByIDRequest{
		Id: "1",
	})
	if err != nil {
		log.Fatalf("could not find user: %v", err)
	}
	log.Printf("User: %v", r.User)
}
