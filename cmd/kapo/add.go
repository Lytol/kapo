package main

import (
	"fmt"
	"log"

	"github.com/Lytol/kapo"
	"github.com/spf13/cobra"
)

var recipient string

func init() {
	addCmd.Flags().StringVarP(&recipient, "recipient", "r", "", "Recipient of transaction")
	addCmd.MarkFlagRequired("recipient")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a transaction to the blockchain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bc, err := kapo.NewBlockchain()
		if err != nil {
			log.Fatal(err)
		}

		data := []byte(args[0])

		tx := kapo.NewTransaction(recipient, data)

		bc.AddBlock([]*kapo.Transaction{tx})

		fmt.Println(tx)
	},
}
