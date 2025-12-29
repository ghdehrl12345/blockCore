package main

import (
	"fmt"

	"github.com/ghdehrl12345/blockCore/core"
)

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("first block")
	bc.AddBlock("second block")
	bc.AddBlock("third block")

	for i, block := range bc.Blocks {
		fmt.Printf("========== Block %d ==========\n", i)
		fmt.Printf("Timestamp:     %d\n", block.Timestamp)
		fmt.Printf("Data:          %s\n", block.Data)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash:          %x\n", block.Hash)
		fmt.Println("==============================")
	}
}
