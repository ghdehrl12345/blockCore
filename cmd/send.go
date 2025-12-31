package main

import (
	"fmt"

	"github.com/ghdehrl12345/blockCore/core"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send coins from one address to another",
	Run: func(cmd *cobra.Command, args []string) {
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		amount, _ := cmd.Flags().GetInt("amount")

		if from == "" || to == "" || amount <= 0 {
			fmt.Println("Usage: send --from ADDRESS --to ADDRESS --amount AMOUNT")
			return
		}

		bc := core.NewBlockchain([]byte(from))
		defer bc.Close()

		// 송신자 계정 정보 조회 (Nonce 확인용)
		senderAccount := bc.State.GetAccount([]byte(from))

		// 트랜잭션 생성
		tx := core.NewTransaction(
			[]byte(from),
			[]byte(to),
			amount,
			senderAccount.Nonce,
		)

		// 블록 채굴 (송신자가 채굴자 역할)
		block := bc.MineBlock([]*core.Transaction{tx}, []byte(from))

		fmt.Println("Success!")
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Printf("From: %s (new balance: %d)\n", from, bc.GetBalance([]byte(from)))
		fmt.Printf("To: %s (new balance: %d)\n", to, bc.GetBalance([]byte(to)))
	},
}

func init() {
	sendCmd.Flags().String("from", "", "Source address")
	sendCmd.Flags().String("to", "", "Destination address")
	sendCmd.Flags().Int("amount", 0, "Amount to send")
	rootCmd.AddCommand(sendCmd)
}
