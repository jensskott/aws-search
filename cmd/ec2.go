package cmd

import (
	"fmt"
	"log"

	ec2Client "github.com/jensskott/aws-search/ec2"
	"github.com/spf13/cobra"
)

func init() {
	// Add commands
	RootCmd.AddCommand(ec2)
	ec2.AddCommand(eip)

	// Add flags
	ec2.Flags().StringSlice("filter", []string{}, "Filter resources")

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
		fmt.Println(ec2.Flags().GetStringSlice("filter"))
		client := ec2Client.NewClient("us-west-2")
		eips, err := client.Ec2DescribeEips()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(eips)
	},
}
