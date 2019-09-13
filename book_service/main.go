package main

import (
	"context"
	"log"
	"net"

	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/y-zumi/grpc-go/proto/book"
	"github.com/y-zumi/grpc-go/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	BookStatusLending = "Lending"
)

type BookService struct {
	client user.UsersClient
}

func NewBookService(client user.UsersClient) *BookService {
	return &BookService{
		client: client,
	}
}

func (s *BookService) FindLendingBookByID(ctx context.Context, req *book.FindLendingBookByIDRequest) (*book.FindLendingBookByIDResponse, error) {
	// request UserService
	findByIDRequest := user.FindByIDRequest{
		Id: faker.UUIDDigit(),
	}
	borrower, err := s.client.FindByID(ctx, &findByIDRequest)
	if err != nil {
		return nil, errors.New("user is not found error")
	}

	return &book.FindLendingBookByIDResponse{
		Book: &book.Book{
			Id:     req.Id,
			Title:  faker.Word(),
			Status: BookStatusLending,
		},
		Borrower: borrower.User,
	}, nil
}

func main() {
	// Start listening port
	lis, err := net.Listen("tcp", ":50011")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register BookService to gRPC Server
	s := grpc.NewServer()
	bookService, err := createBookService()
	if err != nil {
		log.Fatalf("did not create book service: %v", err)
	}
	book.RegisterBooksServer(s, bookService)

	// Add grpc.reflection.v1alpha.ServerReflection
	reflection.Register(s)

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func createBookService() (*BookService, error) {
	cli, err := newUserClient()
	if err != nil {
		return nil, errors.Wrap(err, "did not create user client")
	}

	return NewBookService(cli), nil
}

func newUserClient() (user.UsersClient, error) {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "did not connect localhost:5001")
	}

	return user.NewUsersClient(conn), nil
}
