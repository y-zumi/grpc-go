package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/y-zumi/grpc-go/proto/book"
	"github.com/y-zumi/grpc-go/proto/user"
	"google.golang.org/grpc"
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

func (s *BookService) FindLendingBookByID(ctx context.Context, req book.FindLendingBookByIDRequest) (*book.FindLendingBookByIDResponse, error) {
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
	// Set up a gRPC client
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	usersClient := user.NewUsersClient(conn)

	// Request to gRPC server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	bookService := NewBookService(usersClient)
	resp, err := bookService.FindLendingBookByID(ctx, book.FindLendingBookByIDRequest{
		Id: "123",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Book: %v", resp)
}
