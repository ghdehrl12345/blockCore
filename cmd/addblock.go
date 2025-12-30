package main

import (
	"fmt"

	"github.com/ghdehrl12345/blockCore/core"
	"github.com/spf13/cobra"
)

var addBlockCmd = &cobra.Command{
	Use:   "addblock [data]",
	Short: "Add a new block to the blockchain",
	Args:  cobra.ExactArgs(1), // 정확히 1개의 인자 필요
	Run: func(cmd *cobra.Command, args []string) {
		bc := core.NewBlockchain()
		defer bc.Close()

		block := bc.AddBlock(args[0])
		fmt.Printf("Block added!\n")
		fmt.Printf("Hash: %x\n", block.Hash)
	},
}

func init() {
	rootCmd.AddCommand(addBlockCmd)
}
