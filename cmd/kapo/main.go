package main

import (
	"github.com/spf13/cobra"
)

const (
	Version = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use: "kapo",
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(headCmd)
}

func main() {
	rootCmd.Execute()
}
