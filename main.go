package main

import (
	"fmt"

	"github.com/blockchain/krygz/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Pedro")
	bc.AddBlock("Send 2 more BTC to Pedro")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
