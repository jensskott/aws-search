package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd base command for cli interface
var RootCmd = &cobra.Command{
	Use:   "aws-search",
	Short: "Search for aws resources",
	Long: `Searching aws accounts and regions for regions
			over your whole account.
			Uses the standard aws config or any other
	        means of accessing amazon.`,
}
