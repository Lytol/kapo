package main

import (
	"fmt"
	"log"

	"github.com/Lytol/kapo"
	"github.com/spf13/cobra"
)

var headCmd = &cobra.Command{
	Use:   "head",
	Short: "Print the current head of the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		bc, err := kapo.NewBlockchain()
		if err != nil {
			log.Fatal(err)
		}

		head, err := bc.Head()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%x\n", head)
	},
}
