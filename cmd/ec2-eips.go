package cmd

import (
	"log"

	"os"

	ec2Client "github.com/jensskott/aws-search/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Get elastic ips from ec2 for all regions
var eip = &cobra.Command{
	Use:   "eip",
	Short: "Use to list eip resources",
	Long:  "Use to list eips and apply filters to search",
	Run: func(cmd *cobra.Command, args []string) {
		// Create slice for data
		var eipData [][]string

		// Range over each ec2 region
		for _, r := range Regions {
			// Create a new client for region
			client := ec2Client.NewClient(r)

			// Get eip data
			eips, err := client.Ec2DescribeEips(Filter)
			if err != nil {
				log.Fatal(err)
			}

			// Make sure you append only if you get a respons
			if len(eips) != 0 {
				// Add the data to the slice for data to printout
				for _, d := range eips {
					eipData = append(eipData, []string{*d.PublicIp, *d.PrivateIpAddress, *d.InstanceId, *d.NetworkInterfaceId, r})
				}
			}

		}
		printTable(eipData)
	},
}

func printTable(data [][]string) {
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
}
