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
var igw = &cobra.Command{
	Use:   "igw",
	Short: "Use to list internet gateway resources",
	Long: `Use to list internet gateway and apply filters to search

Available filters are:

attachment.state - The current state of the attachment between
the gateway and the VPC (available ). Present only if a VPC is attached.

attachment.vpc-id - The ID of an attached VPC.

internet-gateway-id - The ID of the Internet gateway.

tag :key =*value* - The key/value combination of a tag assigned to the resource.
Specify the key of the tag in the filter name and the value of the tag in the filter value.
For example, for the tag Purpose=X, specify tag:Purpose for the filter name and X for the filter value.

tag-key - The key of a tag assigned to the resource.
This filter is independent of the tag-value filter. For example,
if you use both the filter "tag-key=Purpose" and the filter "tag-value=X",
you get any resources assigned both the tag key Purpose (regardless of what the tag's value is),
and the tag value X (regardless of what the tag's key is). If you want to list only resources where
Purpose is X, see the tag :key =*value* filter.

tag-value - The value of a tag assigned to the resource. This filter is independent of the tag-key filter.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create slice for data
		var data [][]string
		var rawData []*ec2.InternetGateway

		// Range over each ec2 region
		for _, r := range Regions {
			// Create a new client for region
			client := ec2Client.NewClient(r)

			// Get eip data
			resp, err := client.Ec2DescribeInternetGateways(Filter)
			if err != nil {
				log.Fatal(err)
			}

			// Make sure you append only if you get a respons
			if len(resp) != 0 {
				// Add the data to the slice for data to printout
				for _, d := range resp {
					if Raw == false {
						data = append(data, []string{*d.InternetGatewayId, r})
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
			table.SetHeader([]string{"Internet Gateway ID", "Region"})

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
