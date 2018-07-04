package main

import (
	"abbas/blockchain/chain"
	"abbas/blockchain/proto"
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type Server struct {
	chain *chainpkg.BlockChain
}

func (s *Server) AddBlock(ctx context.Context, req *protopkg.BlockRequest) (*protopkg.BlockResponce, error) {
	block := s.chain.AppendBlock(req.Data)

	return &protopkg.BlockResponce{
		Hash: block.Hash,
	}, nil
}

func (s *Server) GetChain(ctx context.Context, req *protopkg.ChainRequest) (*protopkg.ChainResponce, error) {
	res := new(protopkg.ChainResponce)
	for _, block := range s.chain.Blocks {

		res.Blocks = append(res.Blocks, &protopkg.Block{
			PrvHash: block.PrvHash,
			Data:    block.Data,
			Hash:    block.Hash,
		})

	}
	return res, nil
}

func main() {
	listener, errl := net.Listen("tcp", ":8081")

	if errl != nil {
		log.Fatalf("unable to listen on port : %v", errl)
	}

	srv := grpc.NewServer()
	server := &Server{chainpkg.MakeBlockChain()}

	protopkg.RegisterBlockChainServer(srv, server)

	srv.Serve(listener)
}
