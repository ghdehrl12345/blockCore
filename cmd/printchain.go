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
		// 임시 주소로 체인 로드 (읽기 전용)
		bc := core.NewBlockchain([]byte("temp"))
		defer bc.Close()

		iter := bc.NewIterator()
		for {
			block := iter.Next()
			if block == nil {
				break
			}

			fmt.Printf("========== Block ==========\n")
			fmt.Printf("Timestamp:     %d\n", block.Timestamp)
			fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
			fmt.Printf("Hash:          %x\n", block.Hash)
			fmt.Printf("Nonce:         %d\n", block.Nonce)
			fmt.Printf("Transactions:  %d\n", len(block.Transactions))

			for i, tx := range block.Transactions {
				fmt.Printf("  [TX %d]\n", i)
				fmt.Printf("    ID:     %x\n", tx.ID)
				fmt.Printf("    From:   %s\n", tx.From)
				fmt.Printf("    To:     %s\n", tx.To)
				fmt.Printf("    Amount: %d\n", tx.Amount)
			}

			pow := core.NewProofOfWork(block)
			fmt.Printf("Valid PoW:     %v\n", pow.Validate())
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(printChainCmd)
}
