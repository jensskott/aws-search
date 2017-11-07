package cmd

import (
	"fmt"

	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
	ec2Client "github.com/jensskott/aws-search/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Get elastic ips from ec2 for all regions
var ami = &cobra.Command{
	Use:   "ami",
	Short: "Use to list ami resources",
	Long: `Use to list amis and apply filters to search
		   Available filters are:

	       architecture - The image architecture (i386 | x86_64 ).

	       description - The description of the image (provided during image creation).

		   image-id - The ID of the image.

	       name - The name of the AMI (provided during image creation).

	       owner-id - The AWS account ID of the image owner.

		   platform - The platform. To only list Windows-based AMIs, use windows .

		   state - The state of the image (available | pending | failed ).

		   tag :key =*value* - The key/value combination of a tag assigned to the resource.
		   Specify the key of the tag in the filter name and the value of the tag in the filter value.
		   For example, for the tag Purpose=X, specify tag:Purpose for the filter name and X for the filter value.

	       tag-key - The key of a tag assigned to the resource. This filter is independent of the tag-value filter.
	       For example, if you use both the filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
		   assigned both the tag key Purpose (regardless of what the tag's value is), and the tag value X (regardless of what the tag's key is).
		   If you want to list only resources where Purpose is X, see the tag :key =*value* filter.

		   tag-value - The value of a tag assigned to the resource. This filter is independent of the tag-key filter.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create slice for data
		var data [][]string
		var rawData []*ec2.Image

		// Range over each ec2 region
		for _, r := range Regions {
			// Create a new client for region
			client := ec2Client.NewClient(r)

			// Get eip data
			resp, err := client.Ec2DescribeImages(Filter)
			if err != nil {

			}

			// Make sure you append only if you get a respons
			if len(resp) != 0 {
				// Add the data to the slice for data to printout
				for _, d := range resp {
					if Raw == false {
						data = append(data, []string{*d.Name, *d.ImageId, *d.CreationDate, r})
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
			table.SetHeader([]string{"Name", "Image ID", "Creation Date", "Region"})

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
