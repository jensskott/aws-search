package ec2

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (e *Ec2Implementation) Ec2DescribeInstances(f []string) ([]*ec2.Instance, error) {
	var filter []*ec2.Filter
	var returnResp []*ec2.Instance
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

	params := &ec2.DescribeInstancesInput{
		Filters: filter,
	}

	// Describe all describe in the region
	resp, err := e.Svc.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			returnResp = append(returnResp, inst)
		}
	}
	return returnResp, nil
}
