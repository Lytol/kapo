package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var address string

func init() {
	newTransactionCmd.Flags().StringVarP(&address, "address", "a", "", "Account address for transaction")
	newTransactionCmd.MarkFlagRequired("address")
	transactionCmd.AddCommand(newTransactionCmd)
	rootCmd.AddCommand(transactionCmd)
}

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Commands related to transactions",
}

var newTransactionCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new transaction",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Println("TODO")
	},
}
