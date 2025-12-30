package main

import (
	"fmt"

	"github.com/ghdehrl12345/blockCore/core"
	"github.com/spf13/cobra"
)

var printChainCmd = &cobra.Command{
	Use:   "printchain",
	Short: "Print all blocks in the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		bc := core.NewBlockchain()
		defer bc.Close()

		iter := bc.NewIterator()
		for {
			block := iter.Next()
			if block == nil {
				break
			}

			fmt.Printf("========== Block ==========\n")
			fmt.Printf("Timestamp:     %d\n", block.Timestamp)
			fmt.Printf("Data:          %s\n", block.Data)
			fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
			fmt.Printf("Hash:          %x\n", block.Hash)
			fmt.Printf("Nonce:         %d\n", block.Nonce)

			pow := core.NewProofOfWork(block)
			fmt.Printf("Valid PoW:     %v\n", pow.Validate())
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(printChainCmd)
}
