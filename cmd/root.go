package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "aws-search",
	Short: "Search for aws resources",
	Long: `Searching aws accounts and regions for regions
			over your whole account.
			Uses the standard aws config or any other
	        means of accessing amazon.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Uset -h, --help to get help how to use the application")
	},
}
