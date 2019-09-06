## Overview
1. start gRPC server
    ```bash
    $ go run server/main.go
    ```
1. call gRPC server via gRPC client
    ```bash
   $ go run client/main.go
    ```
1. get response from gRPC server
    ```bash
   2019/09/06 23:15:16 User: id:"1" name:"Sample"
    ```

## When you use grpcurl
1. call gRPC server with grpcurl
    ```bash
    $ echo '{"id": "1"}' | grpcurl -k call localhost:5001 Users.FindByID | jq .
    ```
1. get response from gRPC server
    ```json
    {
      "user": {
        "id": "1",
        "name": "Sample"
      }
    }
    ```