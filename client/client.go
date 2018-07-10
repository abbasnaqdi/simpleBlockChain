package main

import (
	"abbas/blockchain/proto"
	"flag"
	"log"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

var client protopkg.BlockChainClient

func main() {
	start := flag.Bool("start", false, "start mining block")
	add := flag.Bool("add", false, "add block to chain")
	list := flag.Bool("list", false, "get blockchain")

	flag.Parse()

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = protopkg.NewBlockChainClient(conn)

	if *start {
		startMining()
	}

	if *add {
		addBlock()
	}

	if *list {
		getBlockChain()
	}
}

func startMining() {

	for {

		block, addErr := client.AddBlock(context.Background(), &protopkg.BlockRequest{
			Data: time.Now().String(),
		})

		if addErr != nil {
			log.Fatalf("unable to add block : %v", addErr)
		}

		log.Printf("new block hash -> %s \n", block.Hash)

		time.Sleep(1 * time.Second)
	}
}

func addBlock() {
	block, addErr := client.AddBlock(context.Background(), &protopkg.BlockRequest{
		Data: time.Now().String(),
	})
	if addErr != nil {
		log.Fatalf("unable to add block : %v", addErr)
	}
	log.Printf("new block hash -> %s\n", block.Hash)
}

func getBlockChain() {
	blockchain, getErr := client.GetChain(context.Background(), &protopkg.ChainRequest{})
	if getErr != nil {
		log.Fatalf("unable to get blockchainm: %v", getErr)
	}

	log.Println("blocks:")
	for _, b := range blockchain.Blocks {
		log.Printf("prev hash: %s, data %s, , hash: %s \n", b.PrvHash, b.Hash, b.Data)
	}
}
