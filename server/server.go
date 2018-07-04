package server

import (
	"abbas/blockchain/proto"
	"context"
	"crypto/sha1"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

type BlockChain struct {
	blockchain proto.ChainResponce
}

func AddBlocks(context.Context, *BlockRequest) (*BlockResponce, error) {

}

func GetChain(context.Context, *ChainRequest) (*ChainResponce, error) {

}

func main() {
	listener, errl := net.Listen("tcp", ":8081")

	if errl != nil {
		log.Fatalf("unable to listen on port : %v", errl)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServiceServer(srv, &BlockChain{})
	srv.Serve(listener)
}

func makeBlock(prvHash string) proto.Block {
	data := rand.Intn(10000)
	hash := toHash(string(data))

	return proto.Block{PrvHash: prvHash, Data: string(data), Hash: hash}
}

func toHash(value string) string {
	h := sha1.New()
	h.Write([]byte(value))
	return string(h.Sum(nil))
}
