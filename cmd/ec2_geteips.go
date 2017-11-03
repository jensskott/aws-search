package cmd

import (
	"fmt"
	"log"

	"github.com/jensskott/aws-search/ec2"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(getEipsCmd)
}

var getEipsCmd = &cobra.Command{
	Use:   "eip",
	Short: "Get all elastic ips from your account",
	Long:  `List all elastic ips from all regions in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		client := ec2.NewClient("us-west-2")
		eips, err := client.Ec2DescribeEips()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(eips)
	},
}
