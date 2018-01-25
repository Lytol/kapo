package main

import (
	"github.com/spf13/cobra"
)

const (
	Version = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "kapo",
	Short: "Kapo is a minimal distributed ledger using a blockchain",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(showCmd)
}

func main() {
	rootCmd.Execute()
}
