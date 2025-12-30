package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blockcore",
	Short: "A simple blockchain implementation in Go",
	Long:  "BlockCore is a blockchain built from scratch for learning purposes.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
