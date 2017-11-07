package cmd

import (
	ec2Client "github.com/jensskott/aws-search/ec2"
	"github.com/spf13/cobra"
)

// Filter for global access
var Filter []string

// Regions for global access
var Regions []string

// Raw output
var Raw bool

func init() {
	// Add commands
	RootCmd.AddCommand(ec2cmd)
	ec2cmd.AddCommand(eip)
	ec2cmd.AddCommand(cgw)

	// Add flags
	RootCmd.PersistentFlags().StringSliceVarP(&Filter, "filter", "f", nil, "Filter resources in aws")
	RootCmd.PersistentFlags().BoolVarP(&Raw, "raw", "r", false, "To print out raw json from response")

	// Get ec2 regions dynamicly
	client := ec2Client.NewClient("eu-west-1")
	Regions, _ = client.Ec2GetRegions()
}

// Ec2 subcommand for ec2 resources
var ec2cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Use to list ec2 resources",
	Long:  "Subcommand to access the ec2 resources and list them",
}
