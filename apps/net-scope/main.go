package main

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "net-scope",
	Short: "Network visibility for distributed services.",
	Long: "net-scope is a CLI tool for exploring and debugging network topology in the IDP.",
}

var scanCmd = &cobra.Command {
	Use: "scan",
	Short: "scan and display network topology",
	Run: func( cmd *cobra.Command, args[] string) {
		fmt.Println("scanning network topology...")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
