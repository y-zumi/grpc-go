package main

import (
	"context"
	"log"
	"net"

	"github.com/bxcodec/faker/v3"
	"github.com/y-zumi/grpc-go/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// UserService presents pb.UsersService
type UserService struct{}

// FindByID find user by user id
func (s *UserService) FindByID(ctx context.Context, req *user.FindByIDRequest) (*user.FindByIDResponse, error) {
	return &user.FindByIDResponse{
		User: &user.User{
			Id:   req.Id,
			Name: faker.Name(),
		},
	}, nil
}

func main() {
	// Start listening port
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register UsersServer to gRPC Server
	s := grpc.NewServer()
	user.RegisterUsersServer(s, &UserService{})

	// Add grpc.reflection.v1alpha.ServerReflection
	reflection.Register(s)

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
