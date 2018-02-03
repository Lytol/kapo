package main

import (
	"fmt"
	"os"

	"github.com/Lytol/kapo"
	"github.com/spf13/cobra"
)

var address string
var privateKey string

func init() {
	newTransactionCmd.Flags().StringVarP(&address, "address", "a", "", "Account address for transaction")
	newTransactionCmd.Flags().StringVarP(&privateKey, "private", "P", "", "Private key for transaction signing")

	newTransactionCmd.MarkFlagRequired("address")
	newTransactionCmd.MarkFlagRequired("private")

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
		var err error

		addr, err := kapo.HexToAddress(address)
		if err != nil {
			fmt.Printf("%s is not a valid address\n", address)
			os.Exit(1)
		}

		priv, err := kapo.HexToPrivateKey(privateKey)
		if err != nil {
			fmt.Printf("%s is not a valid private key\n", privateKey)
			os.Exit(1)
		}

		tx := kapo.NewTransaction(addr, []byte(args[0]))

		err = tx.Sign(priv)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(tx)
	},
}
