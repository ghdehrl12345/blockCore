package main

import (
	"fmt"

	"github.com/ghdehrl12345/blockCore/core"
	"github.com/spf13/cobra"
)

var getBalanceCmd = &cobra.Command{
	Use:   "getbalance [address]",
	Short: "Get balance of an address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := []byte(args[0])

		bc := core.NewBlockchain(address) // 기존 체인 로드
		defer bc.Close()

		balance := bc.GetBalance(address)
		fmt.Printf("Balance of %s: %d\n", args[0], balance)
	},
}

func init() {
	rootCmd.AddCommand(getBalanceCmd)
}
