package main

import (
	"fmt"
	"strconv"

	"github.com/jorgelabadev/blockchain-learning/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.Blocks {

		fmt.Printf("Previus Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)

		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
