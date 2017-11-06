package ec2

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (e *Ec2Implementation) Ec2DescribeEips(f []string) ([]interface{}, error) {
	// Create interfaces for returns
	var dataSlice []interface{}
	var data interface{}

	// Create filter from slice
	var filter []*ec2.Filter
	if f != nil {
		for _, i := range f {
			s := strings.Split(i, " ")
			x := &ec2.Filter{
				Name:   aws.String(s[0]),
				Values: []*string{aws.String(s[1])},
			}
			filter = append(filter, x)
		}
	}

	params := &ec2.DescribeAddressesInput{
		Filters: filter,
	}

	// Describe all describe in the region
	resp, err := e.Svc.DescribeAddresses(params)
	if err != nil {
		return nil, err
	}

	// Add all data from the respons to the interface
	for _, a := range resp.Addresses {
		data = a
		dataSlice = append(dataSlice, data)
	}

	return dataSlice, nil
}
