package main

import (
	"context"
	"log"
	"net"

	pb "github.com/y-zumi/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// UserService presents pb.UsersService
type UserService struct{}

// FindByID find user by user id
func (s *UserService) FindByID(ctx context.Context, req *pb.FindByIDRequest) (*pb.FindByIDResponse, error) {
	return &pb.FindByIDResponse{
		User: &pb.User{
			Id:   req.Id,
			Name: "Sample",
		},
	}, nil
}

func main() {
	// Start listening port
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register UsersServer to gRPC Server
	s := grpc.NewServer()
	pb.RegisterUsersServer(s, &UserService{})

	// Add grpc.reflection.v1alpha.ServerReflection
	reflection.Register(s)

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
