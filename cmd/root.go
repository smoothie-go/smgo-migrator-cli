/*
Copyright © 2025 smoothie-go <admin@hzqki.me>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smgo-migrator-cli",
	Short: "cli program that migrates a smoothie-rs config to smoothie-go",
	Long: `smgo-migrator-cli is a cli program that migrates a smoothie-rs config to smoothie-go.
Usage:
  smgo-migrator-cli -h | --help; Show this message
  smgo-migrator-cli migrate <input filename> <output filename>; Migrate a smoothie-rs config
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
