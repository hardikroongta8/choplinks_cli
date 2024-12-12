package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new short URL",
	Long:  "Create a new short URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Shortened URL: %s.\n\n", createNewURL(args[0]))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
