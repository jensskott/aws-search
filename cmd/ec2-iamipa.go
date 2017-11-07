package cmd

import (
	"fmt"
	"log"

	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
	ec2Client "github.com/jensskott/aws-search/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Get elastic ips from ec2 for all regions
var iamipa = &cobra.Command{
	Use:   "iamipa",
	Short: "Use to list iam instance profice associations resources",
	Long: `Use to list iam instance profice associations and apply filters to search
		   Available filters are:
		   
		   instance-id - The ID of the instance.
		   
		   state - The state of the association 
		   (associating | associated | disassociating | disassociated ).`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create slice for data
		var data [][]string
		var rawData []*ec2.IamInstanceProfileAssociation

		// Range over each ec2 region
		for _, r := range Regions {
			// Create a new client for region
			client := ec2Client.NewClient(r)

			// Get eip data
			resp, err := client.Ec2DescribeIamInstanceProfileAssociations(Filter)
			if err != nil {
				log.Fatal(err)
			}

			// Make sure you append only if you get a respons
			if len(resp) != 0 {
				// Add the data to the slice for data to printout
				for _, d := range resp {
					if Raw == false {
						data = append(data, []string{*d.IamInstanceProfile.Arn, *d.InstanceId, *d.State, r})
					} else {
						rawData = append(rawData, d)
					}
				}
			}

		}
		if Raw == false {
			// Write to std out
			table := tablewriter.NewWriter(os.Stdout)

			// Set the table header
			table.SetHeader([]string{"Instance Profile Arn", "Instance", "State", "Region"})

			// Append all data to table
			for _, d := range data {
				table.Append(d)
			}

			// Writeout table
			table.Render()
		} else {
			fmt.Println(rawData)
		}
	},
}
