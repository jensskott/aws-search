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
var cgw = &cobra.Command{
	Use:   "cgw",
	Short: "Use to list customer gateway resources",
	Long:  "Use to list customer gateway and apply filters to search",
	Run: func(cmd *cobra.Command, args []string) {
		// Create slice for data
		var data [][]string
		var rawData []*ec2.CustomerGateway

		// Range over each ec2 region
		for _, r := range Regions {
			// Create a new client for region
			client := ec2Client.NewClient(r)

			// Get eip data
			resp, err := client.Ec2DescribeCustomerGateway(Filter)
			if err != nil {
				log.Fatal(err)
			}

			// Make sure you append only if you get a respons
			if len(resp) != 0 {
				// Add the data to the slice for data to printout
				for _, d := range resp {
					if Raw == false {
						data = append(data, []string{*d.CustomerGatewayId, *d.IpAddress, *d.State, r})
					} else {
						rawData = append(rawData, d)
					}
				}
			}

		}
		if Raw == false {
			table := tablewriter.NewWriter(os.Stdout)

			// Set the table header
			table.SetHeader([]string{"Gateway ID", "IP address", "State", "Region"})

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
