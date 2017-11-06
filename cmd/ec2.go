package cmd

import (
	"fmt"
	"log"

	ec2Client "github.com/jensskott/aws-search/ec2"
	"github.com/spf13/cobra"
)

var Filter []string
var Regions []string

func init() {
	// Add commands
	RootCmd.AddCommand(ec2)
	ec2.AddCommand(eip)
	// Add flags
	RootCmd.PersistentFlags().StringSliceVarP(&Filter, "filter", "f", nil, "Filter resources in aws")

	client := ec2Client.NewClient("eu-west-1")
	Regions, _ = client.Ec2GetRegions()
}

var ec2 = &cobra.Command{
	Use:   "ec2",
	Short: "Use to list ec2 resources",
	Long:  "Subcommand to access the ec2 resources and list them",
}

var eip = &cobra.Command{
	Use:   "eip",
	Short: "Use to list eip resources",
	Long:  "Use to list eips and apply filters to search",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Regions)
		client := ec2Client.NewClient("us-west-2")
		eips, err := client.Ec2DescribeEips(Filter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(eips)
	},
}
