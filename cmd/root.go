package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "chop",
	Short: "ChopLinks is a cli tool for shortening URLs!",
	Long:  "ChopLinks is a cli tool for shortening URLs!",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Chop: '%s'\n", err)
		os.Exit(1)
	}
}
