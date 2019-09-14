FROM golang:1.13.0 AS builder
WORKDIR /go/src/github.com/y-zumi/grpc-go
COPY . .
RUN make build-user

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/y-zumi/grpc-go/bin/user /usr/local/bin
ENTRYPOINT ["/usr/local/bin/user"]