## How to deploy to GKE
I wrote blog how to deploy to GKE.  
https://y-zumi.hatenablog.com/entry/2019/09/13/234018

## Overview
1. start User gRPC server
    ```bash
    $ go run user_server/main.go
    ```
1. start Book gRPC server
    ```bash
   $ go run book_server/main.go
    ```
1. access to Book gRPC server with [grpcurl](https://github.com/kazegusuri/grpcurl)
    ```bash
   $ echo '{"id": "j3j32dj0129"}' | grpcurl -k call localhost:50011 book.Books.FindLendingBookByID | jq .
   {
     "book": {
       "id": "j3j32dj0129",
       "title": "commodi",
       "status": "Lending"
     },
     "borrower": {
       "id": "179a14d146514c92aeb519be0c6a4cc6",
       "name": "Miss Maeve Larson"
     }
   }
    ```
