package main

import (
	"fmt"
	"log"

	"github.com/Lytol/kapo"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(newAccountCmd)
	rootCmd.AddCommand(accountCmd)
}

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Commands related to accounts",
}

var newAccountCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new account",
	Run: func(cmd *cobra.Command, args []string) {
		account, err := kapo.NewAccount()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(account)
	},
}
