package ec2

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Ec2DescribeEips to get elastic ips from a region
func (e *Ec2Implementation) Ec2DescribeEips(f []string) ([]*ec2.Address, error) {
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

	return resp.Addresses, nil
}
