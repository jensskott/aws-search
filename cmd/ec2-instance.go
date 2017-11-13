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
var instance = &cobra.Command{
	Use:   "instance",
	Short: "Use to list instance resources",
	Long: `Use to list instance and apply filters to search

Available filters are:

dns-name - The public DNS name of the instance.

iam-instance-profile.arn - The instance profile associated with the instance. Specified as an ARN.

image-id - The ID of the image used to launch the instance.

instance-id - The ID of the instance.

instance-state-name - The state of the instance (pending | running | shutting-down | terminated | stopping | stopped ).

instance-type - The type of instance (for example, t2.micro ).

instance.group-id - The ID of the security group for the instance.

instance.group-name - The name of the security group for the instance.

ip-address - The public IPv4 address of the instance.

private-dns-name - The private IPv4 DNS name of the instance.

private-ip-address - The private IPv4 address of the instance.

subnet-id - The ID of the subnet for the instance.

tag :key =*value* - The key/value combination of a tag assigned to the resource.
Specify the key of the tag in the filter name and the value of the tag in the filter value.
For example, for the tag Purpose=X, specify tag:Purpose for the filter name and X for the filter value.

tag-key - The key of a tag assigned to the resource. This filter is independent of the tag-value filter.
For example, if you use both the filter "tag-key=Purpose" and the filter "tag-value=X", you get any resources
assigned both the tag key Purpose (regardless of what the tag's value is), and the tag value X
(regardless of the tag's key). If you want to list only resources where Purpose is X, see the tag :key =*value* filter.

tag-value - The value of a tag assigned to the resource. This filter is independent of the tag-key filter.

vpc-id - The ID of the VPC that the instance is running in.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create slice for data
		var data [][]string
		var rawData []*ec2.Instance

		// Range over each ec2 region
		for _, r := range Regions {
			// Create a new client for region
			client := ec2Client.NewClient(r)

			// Get eip data
			resp, err := client.Ec2DescribeInstances(Filter)
			if err != nil {
				log.Fatal(err)
			}

			// Make sure you append only if you get a respons
			if len(resp) != 0 {
				// Add the data to the slice for data to printout
				for _, d := range resp {
					if Raw == false {
						data = append(data, []string{*d.InstanceId, *d.IamInstanceProfile.Arn, *d.PrivateIpAddress, *d.PublicIpAddress, r})
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
			table.SetHeader([]string{"Instance ID", "IAM profile", "Private IP", "Public IP", "Region"})

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
