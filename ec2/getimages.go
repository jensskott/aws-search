package ec2

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (e *Ec2Implementation) Ec2DescribeImages(f []string) ([]*ec2.Image, error) {
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
	// #TODO: Use value from account number running as filter value
	// Filter out only own ami images
	filter = append(filter, &ec2.Filter{
		Name:   aws.String("owner-id"),
		Values: []*string{aws.String("962744915737")},
	})

	params := &ec2.DescribeImagesInput{
		Filters: filter,
	}
	// Describe all describe in the region
	resp, err := e.Svc.DescribeImages(params)
	if err != nil {
		return nil, err
	}

	return resp.Images, nil
}
