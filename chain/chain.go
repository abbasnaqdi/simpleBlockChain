package chainpkg

import "crypto/sha1"

type Block struct {
	PrvHash string
	Data    string
	Hash    string
}

type BlockChain struct {
	Blocks []*Block
}

func MakeBlockChain() *BlockChain {
	genesisBlock := makeBlock("genesis hash", "genesis data")
	chain := []*Block{genesisBlock}

	return &BlockChain{chain}
}

func makeBlock(prvHash, data string) *Block {
	block := &Block{
		PrvHash: prvHash,
		Data:    data,
	}

	return block
}

func (chain *BlockChain) AppendBlock(data string) *Block {
	prvBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := makeBlock(prvBlock.toHash(), data)
	chain.Blocks = append(chain.Blocks, newBlock)

	return newBlock
}

func (block *Block) toHash() string {
	h := sha1.New()
	h.Write([]byte(block.Data))

	return string(h.Sum(nil))
}
