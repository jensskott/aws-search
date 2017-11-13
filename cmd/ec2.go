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
	RootCmd.AddCommand(ec2Cmd)
	RootCmd.AddCommand(initCmd)
	ec2Cmd.AddCommand(eip)
	ec2Cmd.AddCommand(cgw)
	ec2Cmd.AddCommand(iamipa)
	ec2Cmd.AddCommand(ami)
	ec2Cmd.AddCommand(instance)
	ec2Cmd.AddCommand(igw)
	ec2Cmd.AddCommand(key)

	// Add flags
	RootCmd.PersistentFlags().StringSliceVarP(&Filter, "filter", "f", nil, "Filter resources in aws")
	RootCmd.PersistentFlags().BoolVarP(&Raw, "raw", "r", false, "To print out raw json from response")

	// Get ec2 regions dynamicly
	client := ec2Client.NewClient("eu-west-1")
	Regions, _ = client.Ec2GetRegions()
}

// Ec2 subcommand for ec2 resources
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Use to list ec2 resources",
	Long:  "Subcommand to access the ec2 resources and list them",
}

// TODO: add init command to get creds for switchrole and get all account numbers into a variable to use in filters
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize aws-search",
	Long:  "Used to initialize aws-search and get assume role credentials for all accounts in the config file",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
