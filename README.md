###### This is a blockchain training project that has been programmed using Protocol Buffers and gRPC.

###### Dependencies :
```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
go get -u golang.org/x/net/context
go get -u github.com/golang/protobuf/proto
 ```

###### Get Started :

start server :
```
go run server/server.go
```

start client :
```
go run client/client.go --start     //start connecting the chains
go run client/client.go --add       //added a block to the chain
go run client/main.go --list        //get all of the generated block chains
```
