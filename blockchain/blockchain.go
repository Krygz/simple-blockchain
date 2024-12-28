package blockchain

import "github.com/blockchain/krygz/block"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	Blocks []*block.Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{block.NewGenesisBlock()}}
}
