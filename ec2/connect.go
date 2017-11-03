package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func NewClient(region string) Ec2Implementation {
	var ec2Client Ec2Implementation
	ec2Client.Session = session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
	ec2Client.Svc = ec2.New(ec2Client.Session)
	return ec2Client
}
