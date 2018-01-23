package main

import (
	"fmt"
	"github.com/Lytol/kapo"
)

func main() {
	bc := kapo.NewBlockchain()

	tx1 := kapo.NewTransaction("brian", []byte("Transaction 1"))
	tx2 := kapo.NewTransaction("adam", []byte("Transaction 2"))
	tx3 := kapo.NewTransaction("gary", []byte("Transaction 3"))

	bc.AddBlock([]*kapo.Transaction{tx1, tx2, tx3})

	for _, block := range bc.Blocks {
		fmt.Printf("ID: %x\n", block.ID)
		fmt.Printf("Prev. ID: %x\n", block.Previous)
		fmt.Printf("Timestamp: %v\n", block.Timestamp)

		for _, tx := range block.Transactions {
			fmt.Printf("  - Transaction: %x | %s | %s\n", tx.ID, tx.Recipient, tx.Data)
		}

		fmt.Println()
	}
}
