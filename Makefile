.PHONY: build
build:
	make build-user-server && \
	make build-book-server

.PHONY: build-user-server
build-user-server:
	go build -o bin/user \
	github.com/y-zumi/grpc-go/user_service/

.PHONY: build-book-server
build-book-server:
	go build -o bin/book \
	github.com/y-zumi/grpc-go/book_service/

