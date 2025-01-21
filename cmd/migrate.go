/*
Copyright © 2025 smoothie-go <admin@hzqki.me>
*/
package cmd

import (
	"fmt"
	mig "github.com/smoothie-go/smoothie-go/migrate"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: migrate,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) {
	if len(cmd.Flags().Args()) != 2 {
		fmt.Println("Usage: smgo-migrator-cli migrate <input filename> <output filename>")
		return
	}
	fmt.Println("Migrating", cmd.Flags().Args()[0])
	Migrated, err := mig.Migrate(cmd.Flags().Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(cmd.Flags().Args()[1], []byte(Migrated), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
