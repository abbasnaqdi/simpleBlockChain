package main

import (
	"abbas/blockchain/chain"
	"abbas/blockchain/proto"
	"flag"
	"log"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

var client protopkg.BlockChainClient

func main() {
	addFlag := flag.Bool("add", false, "Add new block")
	listFlag := flag.Bool("list", false, "List all blocks")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = protopkg.NewBlockChainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	block, addErr := client.AddBlock(context.Background(), &protopkg.BlockRequest{
		Data: time.Now().String(),
	})
	if addErr != nil {
		log.Fatalf("unable to add block: %v", addErr)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockchain() {
	blockchain, getErr := client.GetBlockchain(context.Background(), &protopkg.ChainRequest{})
	if getErr != nil {
		log.Fatalf("unable to get blockchain: %v", getErr)
	}

	log.Println("blocks:")
	for _, b := range blockchain.Blocks {
		log.Printf("hash %s, prev hash: %s, data: %s\n", b.Hash, b.PrevBlockHash, b.Data)
	}
}
