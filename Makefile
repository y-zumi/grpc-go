REVISION := $(shell git rev-parse --short HEAD)
VERSION := commit-$(REVISION)

USER_SERVER_NAME := user-server
BOOK_SERVER_NAME := book-server

GCR_PROJECT := kouzoh-p-y-zumi
REGISTRY := gcr.io/$(GCR_PROJECT)

USER_IMAGE := $(REGISTRY)/$(USER_SERVER_NAME):$(VERSION)
BOOK_IMAGE := $(REGISTRY)/$(BOOK_SERVER_NAME):$(VERSION)

.PHONY: build
build:
	make build-user && \
	make build-book

.PHONY: build-user
build-user:
	CGO_ENABLED=0 go build -o bin/user \
	github.com/y-zumi/grpc-go/user_server/

.PHONY: build-book
build-book:
	CGO_ENABLED=0 go build -o bin/book \
	github.com/y-zumi/grpc-go/book_server/

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: container-user
container-user:
	docker build -t $(USER_IMAGE) \
	 -f docker/user/Dockerfile \
	 .

.PHONY: container-book
container-book:
	docker build -t $(BOOK_IMAGE) \
	 -f docker/book/Dockerfile \
	 .