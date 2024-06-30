package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the application",
	Long:  `A command that prints the version of the application based on the Git tag`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
