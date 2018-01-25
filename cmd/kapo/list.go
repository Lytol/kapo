package main

import (
	"fmt"
	"log"

	"github.com/Lytol/kapo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the entire blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		bc, err := kapo.NewBlockchain()
		if err != nil {
			log.Fatal(err)
		}

		bci := bc.Iterator()

		for {
			block := bci.Next()

			if block == nil {
				break
			}

			fmt.Printf("ID: %x\n", block.ID)
			fmt.Printf("Timestamp: %d\n", block.Timestamp)
			fmt.Printf("Nonce: %x\n", block.Nonce)

			for _, tx := range block.Transactions {
				fmt.Printf("  - TX: %s\n", tx)
			}

			fmt.Println()
		}
	},
}
