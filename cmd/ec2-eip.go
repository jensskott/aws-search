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
var eip = &cobra.Command{
	Use:   "eip",
	Short: "Use to list elastic ips resources",
	Long: `Use to list elastic ips and apply filters to search

Available filters are:
		   
allocation-id - [EC2-VPC] The allocation ID for the address.

association-id - [EC2-VPC] The association ID for the address.

domain - Indicates whether the address is for use in EC2-Classic (standard ) or in a VPC (vpc ).

instance-id - The ID of the instance the address is associated with, if any.

network-interface-id - [EC2-VPC] The ID of the network interface that the address is associated with, if any.

network-interface-owner-id - The AWS account ID of the owner.

private-ip-address - [EC2-VPC] The private IP address associated with the Elastic IP address.

public-ip - The Elastic IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create slice for data
		var data [][]string
		var rawData []*ec2.Address

		// Range over each ec2 region
		for _, r := range Regions {
			// Create a new client for region
			client := ec2Client.NewClient(r)

			// Get eip data
			resp, err := client.Ec2DescribeEips(Filter)
			if err != nil {
				log.Fatal(err)
			}

			// Make sure you append only if you get a respons
			if len(resp) != 0 {
				// Add the data to the slice for data to printout
				for _, d := range resp {
					if Raw == false {
						data = append(data, []string{*d.PublicIp, *d.PrivateIpAddress, *d.InstanceId, *d.NetworkInterfaceId, r})
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
			table.SetHeader([]string{"Public IP", "Private IP", "Instance", "Network Interface", "Region"})

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
