package main

import (
	"fmt"

	"github.com/ghdehrl12345/blockCore/core"
)

func main() {
	bc := core.NewBlockchain()
	defer bc.Close() // 프로그램 종료 시 DB 닫기

	// 모든 블록 출력
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
		fmt.Println()
	}
}
