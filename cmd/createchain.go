package main

import (
	"fmt"

	"github.com/ghdehrl12345/blockCore/core"
	"github.com/spf13/cobra"
)

var createChainCmd = &cobra.Command{
	Use:   "createchain [address]",
	Short: "Create a new blockchain with genesis block",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := []byte(args[0])

		bc := core.NewBlockchain(address)
		defer bc.Close()

		fmt.Println("Blockchain created!")
		fmt.Printf("Genesis block mined. Miner reward: 50 coins to %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(createChainCmd)
}
