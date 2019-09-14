FROM golang:1.13.0-alpine

WORKDIR /go/src/github.com/y-zumi/grpc-go

COPY . /go/src/github.com/y-zumi/grpc-go

EXPOSE "50001"

CMD ["go", "run", "user_server/main.go"]