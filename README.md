#### Blockchain simulation project using gRPC and Protocol Buffers .

#### Setup Packages

 - Install [protoc](https://github.com/google/protobuf/releases) compiler.
 - Install `protoc-gen-go plugin`: `go get -u github.com/golang/protobuf/protoc-gen-go`
 - Define service definition in `.proto` file.
 - Build Go bindings from `.proto` file. `protoc --go_out=plugins=grpc:. proto/blockchain.proto`
 - Install grpc Go package - `go get -u google.golang.org/grpc`.
 - Install context package - `go get -u golang.org/x/net/context`.
 - Install protobuf package - `go get -u github.com/golang/protobuf/proto`

#### 

Start server:
```
go run server/server.go
```

Add block as client:
```
go run client/client.go --add
```

get blockchain as client:
```
go run client/main.go --list
```
